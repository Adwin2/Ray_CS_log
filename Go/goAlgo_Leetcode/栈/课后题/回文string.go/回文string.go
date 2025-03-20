package main

import (
	"../../stack"
)

func Verify(a string, s stack.Stack) bool {
	// index
	for i := 0; i < len(s)/2; i++ {
		s.Push(a[i])
	}

	for i := len(s) - len(s)/2; i < len(s); i++ {
		if s.Peek() == a[i] {
			s.Pop()
		} else {
			return false
		}
	}
	return s.isEmpty()
}

func main() {

}
