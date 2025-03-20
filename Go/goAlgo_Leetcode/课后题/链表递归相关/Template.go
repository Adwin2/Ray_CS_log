package main

import "fmt"

type Node struct {
	Next *Node
	val  int
}

func InitNode(v int) *Node {
	return &Node{
		Next: nil,
		val:  v,
	}
}

func Init(a, b, c int) *Node {
	pa := InitNode(a)
	pb := InitNode(b)
	pc := InitNode(c)
	pa.Next = pb
	pb.Next = pc

	return pa
}

func cnt(p *Node) int {
	if p.Next == nil {
		return 1
	}

	return 1 + cnt(p.Next)
}

func main() {
	p := Init(1, 2, 3)
	fmt.Println(cnt(p))

	fmt.Scanln()
}
