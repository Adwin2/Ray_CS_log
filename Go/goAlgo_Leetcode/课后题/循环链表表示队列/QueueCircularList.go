package main

import "fmt"

type Node struct {
	next *Node
	val  int
}

type ListNode struct {
	rear *Node
}

func NewNode(val int) *Node {
	return &Node{
		next: nil,
		val:  val,
	}
}

func InitListNode() *ListNode {
	p1 := NewNode(1)
	p2 := NewNode(2)
	p3 := NewNode(3)
	p1.next = p2
	p2.next = p3

	ln := &ListNode{p3}
	//ln.rear = p3
	ln.rear.next = p1
	return ln
}

func (ln *ListNode) MakeEmpty() {
	for !ln.IsEmpty() {
		ln.Pop()
	}
}

func (ln *ListNode) IsEmpty() bool {
	return ln.rear == nil
}

// 添加队尾元素
func (ln *ListNode) Push(p *Node) {
	if ln.IsEmpty() && ln.rear == nil {
		ln.rear = p
		return
	}
	tmp := ln.rear.next
	ln.rear.next = p
	p.next = tmp
	ln.rear = p
}

// 弹出队头元素
func (ln *ListNode) Pop() {
	if ln.rear.next == ln.rear {
		ln.rear = nil
		return
	}
	ln.rear.next = ln.rear.next.next
}

func PrintLN(ln *ListNode, n int) error {
	// fmt.Println(ln.rear.val, ln.rear.next.val, ln.rear.next.next.val, ln.rear.next.next.next.val, ln.rear.next.next.next.next.val)
	if ln.rear == nil {
		return fmt.Errorf("ln.rear is nil")
	}
	p := ln.rear
	for i := 0; i < n; i++ {
		fmt.Printf("%d\t", p.next.val)
		p = p.next
	}
	return nil
}

func main() {
	p := InitListNode()
	p.Push(&Node{nil, 4})
	PrintLN(p, 5)
	println()
	p.Pop()
	PrintLN(p, 5)
	println()
	p.Pop()
	PrintLN(p, 5)
	println()
	p.Pop()
	PrintLN(p, 5)
	println()
	p.Pop()
	fmt.Println(PrintLN(p, 5))
	println()
	fmt.Println(p.IsEmpty())
}
