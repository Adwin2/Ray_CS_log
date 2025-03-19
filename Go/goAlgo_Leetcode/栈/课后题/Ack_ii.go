package main

import "sync"

type stack struct {
	elems []int
	lock  sync.Mutex
}

func (s *stack) Push(e ...int) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.elems = append(s.elems, e...)
}

func (s *stack) Pop() int {
	s.lock.Lock()
	defer s.lock.Unlock()
	tmp := s.elems[len(s.elems)-1]
	s.elems = s.elems[:len(s.elems)-1] 
}

func Ack(m, n int) int {
	s := &stack{}
	// if m,n ? build stack
	for m >= 0 ; m --{
		
	}
}

func main() {

}
