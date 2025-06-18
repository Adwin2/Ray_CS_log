package raft

// The file raftapi/raft.go defines the interface that raft must
// expose to servers (or the tester), but see comments below for each
// of these functions for more details.
//
// Make() creates a new raft peer that implements the raft interface.

import (
	"bytes"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"

	"6.5840/labgob"
	"6.5840/labrpc"
	"6.5840/raftapi"
	tester "6.5840/tester1"
)

// Raft server states
type ServerState int

const (
	Follower ServerState = iota
	Candidate
	Leader
)

// Log entry structure
type LogEntry struct {
	Term    int
	Command interface{}
}

// A Go object implementing a single Raft peer.
type Raft struct {
	mu        sync.Mutex          // Lock to protect shared access to this peer's state
	peers     []*labrpc.ClientEnd // RPC end points of all peers
	persister *tester.Persister   // Object to hold this peer's persisted state
	me        int                 // this peer's index into peers[]
	dead      int32               // set by Kill()

	// Your data here (3A, 3B, 3C).
	// Look at the paper's Figure 2 for a description of what
	// state a Raft server must maintain.

	// Persistent state on all servers (updated on stable storage before responding to RPCs)
	currentTerm int        // latest term server has seen (initialized to 0 on first boot, increases monotonically)
	votedFor    int        // candidateId that received vote in current term (or -1 if none)
	log         []LogEntry // log entries; each entry contains command for state machine, and term when entry was received by leader (first index is 1)

	// Volatile state on all servers
	commitIndex int // index of highest log entry known to be committed (initialized to 0, increases monotonically)
	lastApplied int // index of highest log entry applied to state machine (initialized to 0, increases monotonically)

	// Volatile state on leaders (reinitialized after election)
	nextIndex  []int // for each server, index of the next log entry to send to that server (initialized to leader last log index + 1)
	matchIndex []int // for each server, index of highest log entry known to be replicated on server (initialized to 0, increases monotonically)

	// Additional state for implementation
	state           ServerState   // current server state
	electionTimeout time.Time     // when to start next election
	applyCh         chan raftapi.ApplyMsg // channel to send committed entries

	// Snapshot state
	lastIncludedIndex int    // index of last entry in snapshot
	lastIncludedTerm  int    // term of last entry in snapshot
	snapshot          []byte // snapshot data
}

// return currentTerm and whether this server
// believes it is the leader.
func (rf *Raft) GetState() (int, bool) {
	rf.mu.Lock()
	defer rf.mu.Unlock()

	term := rf.currentTerm
	isleader := rf.state == Leader
	return term, isleader
}

// save Raft's persistent state to stable storage,
// where it can later be retrieved after a crash and restart.
// see paper's Figure 2 for a description of what should be persistent.
// before you've implemented snapshots, you should pass nil as the
// second argument to persister.Save().
// after you've implemented snapshots, pass the current snapshot
// (or nil if there's not yet a snapshot).
func (rf *Raft) persist() {
	// Your code here (3C).
	w := new(bytes.Buffer)
	e := labgob.NewEncoder(w)
	e.Encode(rf.currentTerm)
	e.Encode(rf.votedFor)
	e.Encode(rf.log)
	e.Encode(rf.lastIncludedIndex)
	e.Encode(rf.lastIncludedTerm)
	raftstate := w.Bytes()
	rf.persister.Save(raftstate, rf.snapshot)
}

// Persist state and snapshot
func (rf *Raft) persistStateAndSnapshot() {
	w := new(bytes.Buffer)
	e := labgob.NewEncoder(w)
	e.Encode(rf.currentTerm)
	e.Encode(rf.votedFor)
	e.Encode(rf.log)
	e.Encode(rf.lastIncludedIndex)
	e.Encode(rf.lastIncludedTerm)
	raftstate := w.Bytes()
	rf.persister.Save(raftstate, rf.snapshot)
}


// restore previously persisted state.
func (rf *Raft) readPersist(data []byte) {
	if data == nil || len(data) < 1 { // bootstrap without any state?
		return
	}
	// Your code here (3C).
	r := bytes.NewBuffer(data)
	d := labgob.NewDecoder(r)
	var currentTerm int
	var votedFor int
	var log []LogEntry
	var lastIncludedIndex int
	var lastIncludedTerm int
	if d.Decode(&currentTerm) != nil ||
		d.Decode(&votedFor) != nil ||
		d.Decode(&log) != nil ||
		d.Decode(&lastIncludedIndex) != nil ||
		d.Decode(&lastIncludedTerm) != nil {
		// error reading persisted state
	} else {
		rf.currentTerm = currentTerm
		rf.votedFor = votedFor
		rf.log = log
		rf.lastIncludedIndex = lastIncludedIndex
		rf.lastIncludedTerm = lastIncludedTerm
		rf.snapshot = rf.persister.ReadSnapshot()
	}
}

// how many bytes in Raft's persisted log?
func (rf *Raft) PersistBytes() int {
	rf.mu.Lock()
	defer rf.mu.Unlock()
	return rf.persister.RaftStateSize()
}


// the service says it has created a snapshot that has
// all info up to and including index. this means the
// service no longer needs the log through (and including)
// that index. Raft should now trim its log as much as possible.
func (rf *Raft) Snapshot(index int, snapshot []byte) {
	rf.mu.Lock()
	defer rf.mu.Unlock()

	// Don't snapshot if index is not greater than current snapshot
	if index <= rf.lastIncludedIndex {
		return
	}

	// Don't snapshot beyond what we have committed
	if index > rf.commitIndex {
		return
	}

	// Calculate the relative index in the current log
	relativeIndex := index - rf.lastIncludedIndex
	if relativeIndex >= len(rf.log) {
		return
	}

	// Save the term of the entry being snapshotted
	snapshotTerm := rf.log[relativeIndex].Term

	// Create new log starting from the entry after the snapshot
	newLog := make([]LogEntry, 1)
	newLog[0] = LogEntry{Term: snapshotTerm, Command: nil} // Dummy entry at index 0
	if relativeIndex+1 < len(rf.log) {
		newLog = append(newLog, rf.log[relativeIndex+1:]...)
	}

	// Update snapshot state
	rf.lastIncludedIndex = index
	rf.lastIncludedTerm = snapshotTerm
	rf.log = newLog
	rf.snapshot = snapshot

	// Update lastApplied to reflect that entries up to index are applied
	if rf.lastApplied < index {
		rf.lastApplied = index
	}

	// Persist the new state and snapshot
	rf.persistStateAndSnapshot()
}


// RequestVote RPC arguments structure.
// field names must start with capital letters!
type RequestVoteArgs struct {
	Term         int // candidate's term
	CandidateId  int // candidate requesting vote
	LastLogIndex int // index of candidate's last log entry
	LastLogTerm  int // term of candidate's last log entry
}

// RequestVote RPC reply structure.
// field names must start with capital letters!
type RequestVoteReply struct {
	Term        int  // currentTerm, for candidate to update itself
	VoteGranted bool // true means candidate received vote
}

// RequestVote RPC handler.
func (rf *Raft) RequestVote(args *RequestVoteArgs, reply *RequestVoteReply) {
	rf.mu.Lock()
	defer rf.mu.Unlock()

	// Initialize reply
	reply.Term = rf.currentTerm
	reply.VoteGranted = false

	// Rule 1: Reply false if term < currentTerm
	if args.Term < rf.currentTerm {
		return
	}

	// If RPC request or response contains term T > currentTerm:
	// set currentTerm = T, convert to follower
	if args.Term > rf.currentTerm {
		rf.currentTerm = args.Term
		rf.votedFor = -1
		rf.state = Follower
		rf.persist()
		reply.Term = rf.currentTerm
	}

	// Rule 2: If votedFor is null or candidateId, and candidate's log is at
	// least as up-to-date as receiver's log, grant vote
	lastLogIndex := rf.lastIncludedIndex + len(rf.log) - 1
	lastLogTerm := 0
	if len(rf.log) > 1 {
		lastLogTerm = rf.log[len(rf.log)-1].Term
	} else if rf.lastIncludedIndex > 0 {
		lastLogTerm = rf.lastIncludedTerm
	}

	logUpToDate := args.LastLogTerm > lastLogTerm ||
		(args.LastLogTerm == lastLogTerm && args.LastLogIndex >= lastLogIndex)

	if (rf.votedFor == -1 || rf.votedFor == args.CandidateId) && logUpToDate {
		rf.votedFor = args.CandidateId
		rf.persist()
		reply.VoteGranted = true
		rf.resetElectionTimeout()
	}
}

// example code to send a RequestVote RPC to a server.
// server is the index of the target server in rf.peers[].
// expects RPC arguments in args.
// fills in *reply with RPC reply, so caller should
// pass &reply.
// the types of the args and reply passed to Call() must be
// the same as the types of the arguments declared in the
// handler function (including whether they are pointers).
//
// The labrpc package simulates a lossy network, in which servers
// may be unreachable, and in which requests and replies may be lost.
// Call() sends a request and waits for a reply. If a reply arrives
// within a timeout interval, Call() returns true; otherwise
// Call() returns false. Thus Call() may not return for a while.
// A false return can be caused by a dead server, a live server that
// can't be reached, a lost request, or a lost reply.
//
// Call() is guaranteed to return (perhaps after a delay) *except* if the
// handler function on the server side does not return.  Thus there
// is no need to implement your own timeouts around Call().
//
// look at the comments in ../labrpc/labrpc.go for more details.
//
// if you're having trouble getting RPC to work, check that you've
// capitalized all field names in structs passed over RPC, and
// that the caller passes the address of the reply struct with &, not
// the struct itself.
func (rf *Raft) sendRequestVote(server int, args *RequestVoteArgs, reply *RequestVoteReply) bool {
	ok := rf.peers[server].Call("Raft.RequestVote", args, reply)
	return ok
}

// Send AppendEntries RPC
func (rf *Raft) sendAppendEntriesRPC(server int, args *AppendEntriesArgs, reply *AppendEntriesReply) bool {
	ok := rf.peers[server].Call("Raft.AppendEntries", args, reply)
	return ok
}

// InstallSnapshot RPC arguments structure.
type InstallSnapshotArgs struct {
	Term              int    // leader's term
	LeaderId          int    // so follower can redirect clients
	LastIncludedIndex int    // the snapshot replaces all entries up through and including this index
	LastIncludedTerm  int    // term of lastIncludedIndex
	Data              []byte // raw bytes of the snapshot chunk, starting at offset
}

// InstallSnapshot RPC reply structure.
type InstallSnapshotReply struct {
	Term int // currentTerm, for leader to update itself
}

// InstallSnapshot RPC handler.
func (rf *Raft) InstallSnapshot(args *InstallSnapshotArgs, reply *InstallSnapshotReply) {
	rf.mu.Lock()
	defer rf.mu.Unlock()

	reply.Term = rf.currentTerm

	// Rule 1: Reply immediately if term < currentTerm
	if args.Term < rf.currentTerm {
		return
	}

	// If RPC request contains term T > currentTerm: set currentTerm = T, convert to follower
	if args.Term > rf.currentTerm {
		rf.currentTerm = args.Term
		rf.votedFor = -1
		rf.state = Follower
		rf.persist()
		reply.Term = rf.currentTerm
	}

	// Reset election timeout since we heard from leader
	rf.resetElectionTimeout()

	// Rule 6: If existing log entry has same index and term as snapshot's
	// last included entry, retain log entries following it and reply
	if args.LastIncludedIndex <= rf.lastIncludedIndex {
		return // Already have this snapshot or a newer one
	}

	// Save snapshot file, discard any existing or partial snapshot with a smaller index
	rf.snapshot = args.Data
	rf.lastIncludedIndex = args.LastIncludedIndex
	rf.lastIncludedTerm = args.LastIncludedTerm

	// Calculate the current last log index
	currentLastIndex := rf.lastIncludedIndex + len(rf.log) - 1

	// Discard the entire log if snapshot is more recent than our log
	if args.LastIncludedIndex >= currentLastIndex {
		rf.log = make([]LogEntry, 1)
		rf.log[0] = LogEntry{Term: args.LastIncludedTerm, Command: nil}
	} else {
		// Keep log entries after the snapshot
		newLog := make([]LogEntry, 1)
		newLog[0] = LogEntry{Term: args.LastIncludedTerm, Command: nil}

		// Find the relative index where to start keeping entries
		keepFromIndex := args.LastIncludedIndex + 1 - rf.lastIncludedIndex
		if keepFromIndex > 0 && keepFromIndex < len(rf.log) {
			newLog = append(newLog, rf.log[keepFromIndex:]...)
		}
		rf.log = newLog
	}

	rf.commitIndex = max(rf.commitIndex, args.LastIncludedIndex)
	rf.lastApplied = args.LastIncludedIndex

	// Persist state and snapshot
	rf.persistStateAndSnapshot()

	// Send snapshot to application
	msg := raftapi.ApplyMsg{
		SnapshotValid: true,
		Snapshot:      args.Data,
		SnapshotTerm:  args.LastIncludedTerm,
		SnapshotIndex: args.LastIncludedIndex,
	}

	go func() {
		rf.applyCh <- msg
	}()
}

// Send InstallSnapshot RPC
func (rf *Raft) sendInstallSnapshot(server int, args *InstallSnapshotArgs, reply *InstallSnapshotReply) bool {
	ok := rf.peers[server].Call("Raft.InstallSnapshot", args, reply)
	return ok
}

// AppendEntries RPC arguments structure.
type AppendEntriesArgs struct {
	Term         int        // leader's term
	LeaderId     int        // so follower can redirect clients
	PrevLogIndex int        // index of log entry immediately preceding new ones
	PrevLogTerm  int        // term of prevLogIndex entry
	Entries      []LogEntry // log entries to store (empty for heartbeat; may send more than one for efficiency)
	LeaderCommit int        // leader's commitIndex
}

// AppendEntries RPC reply structure.
type AppendEntriesReply struct {
	Term    int  // currentTerm, for leader to update itself
	Success bool // true if follower contained entry matching prevLogIndex and prevLogTerm

	// Optimization for fast backup (binary search optimization)
	XTerm  int // term in the conflicting entry (if any)
	XIndex int // index of first entry with that term (if any)
	XLen   int // log length
}

// AppendEntries RPC handler.
func (rf *Raft) AppendEntries(args *AppendEntriesArgs, reply *AppendEntriesReply) {
	rf.mu.Lock()
	defer rf.mu.Unlock()

	// Initialize reply
	reply.Term = rf.currentTerm
	reply.Success = false

	// Rule 1: Reply false if term < currentTerm
	if args.Term < rf.currentTerm {
		return
	}

	// If RPC request or response contains term T > currentTerm:
	// set currentTerm = T, convert to follower
	if args.Term > rf.currentTerm {
		rf.currentTerm = args.Term
		rf.votedFor = -1
		rf.state = Follower
		rf.persist()
		reply.Term = rf.currentTerm
	}

	// Reset election timeout since we heard from leader
	rf.resetElectionTimeout()

	// Rule 2: Reply false if log doesn't contain an entry at prevLogIndex
	// whose term matches prevLogTerm
	if args.PrevLogIndex > rf.lastIncludedIndex {
		relativeIndex := args.PrevLogIndex - rf.lastIncludedIndex
		if relativeIndex >= len(rf.log) || rf.log[relativeIndex].Term != args.PrevLogTerm {
			// Binary search optimization: provide information for fast backup
			reply.XLen = rf.lastIncludedIndex + len(rf.log)
			if relativeIndex >= len(rf.log) {
				// Case 3: follower's log is too short
				reply.XTerm = -1
				reply.XIndex = rf.lastIncludedIndex + len(rf.log)
			} else {
				// Case 1 or 2: conflicting term
				reply.XTerm = rf.log[relativeIndex].Term
				// Find first index with XTerm
				reply.XIndex = args.PrevLogIndex
				for reply.XIndex > rf.lastIncludedIndex {
					relIdx := reply.XIndex - 1 - rf.lastIncludedIndex
					if relIdx >= 0 && relIdx < len(rf.log) && rf.log[relIdx].Term == reply.XTerm {
						reply.XIndex--
					} else {
						break
					}
				}
			}
			return
		}
	} else if args.PrevLogIndex < rf.lastIncludedIndex {
		// PrevLogIndex is in the snapshot, this is outdated
		reply.XLen = rf.lastIncludedIndex + len(rf.log)
		reply.XTerm = -1
		reply.XIndex = rf.lastIncludedIndex + 1
		return
	} else if args.PrevLogIndex == rf.lastIncludedIndex {
		// PrevLogIndex matches snapshot boundary, check term
		if args.PrevLogTerm != rf.lastIncludedTerm {
			reply.XLen = rf.lastIncludedIndex + len(rf.log)
			reply.XTerm = rf.lastIncludedTerm
			reply.XIndex = rf.lastIncludedIndex
			return
		}
	}

	// Rule 3 & 4: Handle log entries
	if len(args.Entries) > 0 {
		index := args.PrevLogIndex + 1
		logChanged := false

		// Check for conflicts and truncate if necessary
		for i, entry := range args.Entries {
			entryIndex := index + i
			relativeIndex := entryIndex - rf.lastIncludedIndex

			if relativeIndex > 0 && relativeIndex < len(rf.log) {
				if rf.log[relativeIndex].Term != entry.Term {
					// Conflict found, truncate log from this point
					rf.log = rf.log[:relativeIndex]
					logChanged = true
					break
				}
			}
		}

		// Append new entries
		for i, entry := range args.Entries {
			entryIndex := index + i
			relativeIndex := entryIndex - rf.lastIncludedIndex

			if relativeIndex > 0 && relativeIndex >= len(rf.log) {
				rf.log = append(rf.log, entry)
				logChanged = true
			}
		}

		if logChanged {
			rf.persist()
		}
	}

	// Rule 5: If leaderCommit > commitIndex, set commitIndex =
	// min(leaderCommit, index of last new entry)
	if args.LeaderCommit > rf.commitIndex {
		lastNewEntryIndex := rf.lastIncludedIndex + len(rf.log) - 1
		rf.commitIndex = min(args.LeaderCommit, lastNewEntryIndex)
		go rf.applyEntries()
	}

	reply.Success = true
}

// Helper function to reset election timeout
func (rf *Raft) resetElectionTimeout() {
	// Election timeout should be randomized between 150-300ms according to paper
	// But tests require larger timeouts, so we use 300-600ms
	timeout := 300 + rand.Intn(300)
	rf.electionTimeout = time.Now().Add(time.Duration(timeout) * time.Millisecond)
}

// Helper function for min
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Helper function for max
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Update commit index based on majority replication
func (rf *Raft) updateCommitIndex() {
	if rf.state != Leader {
		return
	}

	// Find the highest index that is replicated on majority of servers
	for i := rf.lastIncludedIndex + len(rf.log) - 1; i > rf.commitIndex; i-- {
		relativeIndex := i - rf.lastIncludedIndex
		if relativeIndex > 0 && relativeIndex < len(rf.log) && rf.log[relativeIndex].Term == rf.currentTerm {
			count := 1 // count self
			for j := range rf.peers {
				if j != rf.me && rf.matchIndex[j] >= i {
					count++
				}
			}
			if count > len(rf.peers)/2 {
				rf.commitIndex = i
				go rf.applyEntries()
				break
			}
		}
	}
}

// Apply committed entries to state machine
func (rf *Raft) applyEntries() {
	rf.mu.Lock()

	msgs := []raftapi.ApplyMsg{}

	// Ensure lastApplied is at least at the snapshot boundary
	if rf.lastApplied < rf.lastIncludedIndex {
		rf.lastApplied = rf.lastIncludedIndex
	}

	for rf.lastApplied < rf.commitIndex {
		rf.lastApplied++
		if rf.lastApplied > rf.lastIncludedIndex {
			relativeIndex := rf.lastApplied - rf.lastIncludedIndex
			if relativeIndex > 0 && relativeIndex < len(rf.log) {
				entry := rf.log[relativeIndex]

				msg := raftapi.ApplyMsg{
					CommandValid: true,
					Command:      entry.Command,
					CommandIndex: rf.lastApplied,
				}
				msgs = append(msgs, msg)
			}
		}
	}
	rf.mu.Unlock()

	// Send messages in order without holding lock
	for _, msg := range msgs {
		rf.applyCh <- msg
	}
}


// the service using Raft (e.g. a k/v server) wants to start
// agreement on the next command to be appended to Raft's log. if this
// server isn't the leader, returns false. otherwise start the
// agreement and return immediately. there is no guarantee that this
// command will ever be committed to the Raft log, since the leader
// may fail or lose an election. even if the Raft instance has been killed,
// this function should return gracefully.
//
// the first return value is the index that the command will appear at
// if it's ever committed. the second return value is the current
// term. the third return value is true if this server believes it is
// the leader.
func (rf *Raft) Start(command interface{}) (int, int, bool) {
	rf.mu.Lock()
	defer rf.mu.Unlock()

	index := -1
	term := rf.currentTerm
	isLeader := rf.state == Leader

	if isLeader {
		// Append entry to leader's log
		entry := LogEntry{
			Term:    rf.currentTerm,
			Command: command,
		}
		rf.log = append(rf.log, entry)
		rf.persist()
		index = rf.lastIncludedIndex + len(rf.log) - 1 // This will be the actual index

		// Update nextIndex for this server
		rf.matchIndex[rf.me] = index

		// Start replicating to followers immediately
		go rf.replicateEntries()
	}

	return index, term, isLeader
}

// the tester doesn't halt goroutines created by Raft after each test,
// but it does call the Kill() method. your code can use killed() to
// check whether Kill() has been called. the use of atomic avoids the
// need for a lock.
//
// the issue is that long-running goroutines use memory and may chew
// up CPU time, perhaps causing later tests to fail and generating
// confusing debug output. any goroutine with a long-running loop
// should call killed() to check whether it should stop.
func (rf *Raft) Kill() {
	atomic.StoreInt32(&rf.dead, 1)
	// Your code here, if desired.
}

func (rf *Raft) killed() bool {
	z := atomic.LoadInt32(&rf.dead)
	return z == 1
}

func (rf *Raft) ticker() {
	for rf.killed() == false {
		rf.mu.Lock()

		// Check if a leader election should be started
		if rf.state != Leader && time.Now().After(rf.electionTimeout) {
			rf.startElection()
		}

		rf.mu.Unlock()

		// pause for a random amount of time between 50 and 350
		// milliseconds.
		ms := 50 + (rand.Int63() % 300)
		time.Sleep(time.Duration(ms) * time.Millisecond)
	}
}

// Start a new election
func (rf *Raft) startElection() {
	rf.state = Candidate
	rf.currentTerm++
	rf.votedFor = rf.me
	rf.persist()
	rf.resetElectionTimeout()

	lastLogIndex := rf.lastIncludedIndex + len(rf.log) - 1
	lastLogTerm := 0
	if len(rf.log) > 1 {
		lastLogTerm = rf.log[len(rf.log)-1].Term
	} else if rf.lastIncludedIndex > 0 {
		lastLogTerm = rf.lastIncludedTerm
	}

	args := RequestVoteArgs{
		Term:         rf.currentTerm,
		CandidateId:  rf.me,
		LastLogIndex: lastLogIndex,
		LastLogTerm:  lastLogTerm,
	}

	currentTerm := rf.currentTerm
	votes := 1 // vote for self
	var mu sync.Mutex

	// Send RequestVote RPCs to all other servers
	for i := range rf.peers {
		if i != rf.me {
			go func(server int) {
				reply := RequestVoteReply{}
				if rf.sendRequestVote(server, &args, &reply) {
					rf.mu.Lock()
					defer rf.mu.Unlock()

					// Check if we're still a candidate and in the same term
					if rf.state != Candidate || rf.currentTerm != currentTerm {
						return
					}

					if reply.Term > rf.currentTerm {
						rf.currentTerm = reply.Term
						rf.votedFor = -1
						rf.state = Follower
						rf.persist()
						return
					}

					if reply.VoteGranted {
						mu.Lock()
						votes++
						// Check if we have majority
						if votes > len(rf.peers)/2 && rf.state == Candidate {
							rf.becomeLeader()
						}
						mu.Unlock()
					}
				}
			}(i)
		}
	}
}

// Become leader
func (rf *Raft) becomeLeader() {
	rf.state = Leader

	// Initialize leader state
	rf.nextIndex = make([]int, len(rf.peers))
	rf.matchIndex = make([]int, len(rf.peers))

	for i := range rf.peers {
		rf.nextIndex[i] = rf.lastIncludedIndex + len(rf.log) // Next index to send
		rf.matchIndex[i] = 0                                 // No entries matched initially
	}

	// Send initial heartbeats
	go rf.sendHeartbeats()
}

// Send heartbeats to all followers
func (rf *Raft) sendHeartbeats() {
	for rf.killed() == false {
		rf.mu.Lock()
		if rf.state != Leader {
			rf.mu.Unlock()
			return
		}

		for i := range rf.peers {
			if i != rf.me {
				go rf.sendAppendEntries(i)
			}
		}
		rf.mu.Unlock()

		// Send heartbeats every 100ms (10 times per second as required)
		time.Sleep(100 * time.Millisecond)
	}
}

// Replicate entries to all followers
func (rf *Raft) replicateEntries() {
	rf.mu.Lock()
	if rf.state != Leader {
		rf.mu.Unlock()
		return
	}

	for i := range rf.peers {
		if i != rf.me {
			go rf.sendAppendEntries(i)
		}
	}
	rf.mu.Unlock()
}

// Send AppendEntries RPC to a specific server
func (rf *Raft) sendAppendEntries(server int) {
	rf.mu.Lock()
	if rf.state != Leader {
		rf.mu.Unlock()
		return
	}

	// Check if we need to send snapshot instead
	if rf.nextIndex[server] <= rf.lastIncludedIndex {
		// Send InstallSnapshot RPC
		args := InstallSnapshotArgs{
			Term:              rf.currentTerm,
			LeaderId:          rf.me,
			LastIncludedIndex: rf.lastIncludedIndex,
			LastIncludedTerm:  rf.lastIncludedTerm,
			Data:              rf.snapshot,
		}
		rf.mu.Unlock()

		reply := InstallSnapshotReply{}
		if rf.sendInstallSnapshot(server, &args, &reply) {
			rf.mu.Lock()
			defer rf.mu.Unlock()

			if rf.state != Leader || rf.currentTerm != args.Term {
				return
			}

			if reply.Term > rf.currentTerm {
				rf.currentTerm = reply.Term
				rf.votedFor = -1
				rf.state = Follower
				rf.persist()
				return
			}

			// Update nextIndex and matchIndex after successful snapshot
			rf.nextIndex[server] = rf.lastIncludedIndex + 1
			rf.matchIndex[server] = rf.lastIncludedIndex
		}
		return
	}

	prevLogIndex := rf.nextIndex[server] - 1
	prevLogTerm := 0

	// Calculate relative index for log access
	relativeIndex := prevLogIndex - rf.lastIncludedIndex
	if relativeIndex > 0 && relativeIndex < len(rf.log) {
		prevLogTerm = rf.log[relativeIndex].Term
	} else if prevLogIndex == rf.lastIncludedIndex {
		prevLogTerm = rf.lastIncludedTerm
	}

	// Send entries starting from nextIndex[server]
	entries := []LogEntry{}
	nextRelativeIndex := rf.nextIndex[server] - rf.lastIncludedIndex
	if nextRelativeIndex < len(rf.log) && nextRelativeIndex > 0 {
		entries = rf.log[nextRelativeIndex:]
	}

	args := AppendEntriesArgs{
		Term:         rf.currentTerm,
		LeaderId:     rf.me,
		PrevLogIndex: prevLogIndex,
		PrevLogTerm:  prevLogTerm,
		Entries:      entries,
		LeaderCommit: rf.commitIndex,
	}
	rf.mu.Unlock()

	reply := AppendEntriesReply{}
	if rf.sendAppendEntriesRPC(server, &args, &reply) {
		rf.mu.Lock()
		defer rf.mu.Unlock()

		if rf.state != Leader || rf.currentTerm != args.Term {
			return
		}

		if reply.Term > rf.currentTerm {
			rf.currentTerm = reply.Term
			rf.votedFor = -1
			rf.state = Follower
			rf.persist()
			return
		}

		if reply.Success {
			// Update matchIndex and nextIndex
			rf.matchIndex[server] = args.PrevLogIndex + len(args.Entries)
			rf.nextIndex[server] = rf.matchIndex[server] + 1

			// Check if we can commit more entries
			rf.updateCommitIndex()
		} else {
			// Binary search optimization for fast backup
			if reply.XTerm == -1 {
				// Case 3: follower's log is too short
				rf.nextIndex[server] = reply.XLen
			} else {
				// Case 1 or 2: find last entry with XTerm in leader's log
				lastIndex := -1
				for i := len(rf.log) - 1; i >= 0; i-- {
					if rf.log[i].Term == reply.XTerm {
						lastIndex = i
						break
					}
				}

				if lastIndex == -1 {
					// Case 1: leader doesn't have XTerm
					rf.nextIndex[server] = reply.XIndex
				} else {
					// Case 2: leader has XTerm
					rf.nextIndex[server] = lastIndex + 1
				}
			}

			// Ensure nextIndex doesn't go below 1
			if rf.nextIndex[server] < 1 {
				rf.nextIndex[server] = 1
			}
		}
	}
}

// the service or tester wants to create a Raft server. the ports
// of all the Raft servers (including this one) are in peers[]. this
// server's port is peers[me]. all the servers' peers[] arrays
// have the same order. persister is a place for this server to
// save its persistent state, and also initially holds the most
// recent saved state, if any. applyCh is a channel on which the
// tester or service expects Raft to send ApplyMsg messages.
// Make() must return quickly, so it should start goroutines
// for any long-running work.
func Make(peers []*labrpc.ClientEnd, me int,
	persister *tester.Persister, applyCh chan raftapi.ApplyMsg) raftapi.Raft {
	rf := &Raft{}
	rf.peers = peers
	rf.persister = persister
	rf.me = me
	rf.applyCh = applyCh

	// Initialize Raft state
	rf.currentTerm = 0
	rf.votedFor = -1
	rf.log = make([]LogEntry, 1) // Start with dummy entry at index 0
	rf.log[0] = LogEntry{Term: 0, Command: nil}
	rf.commitIndex = 0
	rf.lastApplied = 0
	rf.lastIncludedIndex = 0
	rf.lastIncludedTerm = 0
	rf.snapshot = nil
	rf.state = Follower
	rf.resetElectionTimeout()

	// initialize from state persisted before a crash
	rf.readPersist(persister.ReadRaftState())

	// start ticker goroutine to start elections
	go rf.ticker()

	return rf
}
