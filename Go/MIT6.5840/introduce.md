> distributed system

#importance:
connect physically separated machines (sharing)
increase capacity through paralielsm
tolerate faults
achieve security/via isolate


*historical context:*
...

#challenges:  (complexity)
many concurrent part
deal with partial failure

#why take 6.824: (now 6.5840)
	hard problem powerful solutions
	used in real world  Active area of research

#course structure:
 lectures: big ideas
 papers: case study

#Labs:(test cases are public(tricky))
 1) mapreduce lib (same as the paper)
 2) replication using Raft 
 3) replicted k/v service
 4) sharded k/v services  /or optional project

#Focus
 infrastructure  Storage Computation Communication(6.829)  ->abstractions

#main topics:
 Fault tolerance ->availability(key is replication) & recoveraniety(..logging and transactions durable storage)
 Consistency between distinct types
 Performance (replication lower it) : throughput & tail latency
 Implementation

#Mapreduce
illustraction Influtial

#context
two persons from Google search engine  multihours Tbs data Web indexing
goal : easy for non-expert



