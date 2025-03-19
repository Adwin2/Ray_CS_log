package main

type ListNode struct {
	Next *ListNode
	val  int
}

// max int
func (l *ListNode) Max() (m int) {
	if l.Next == nil {
		if 
		return 
	}
	return 
}

// cnt
func (l *ListNode) Cnt() int {
	if l == nil {
		return 0
	}
	return 1 + l.Next.Cnt()
}

func main() {
}
