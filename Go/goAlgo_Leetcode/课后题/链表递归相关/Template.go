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

// ① 求链表中的最大整数
func (p *Node) max() int {
	//“归”： 到最后一个节点，返回本值
	if p.Next == nil {
		return p.val
	}
	//“递”： m 是当前p的下一个节点，到最后一个节点时返回值，与栈中所存p.val逐个比较，并最终返回最大值
	m := p.Next.max()
	if m > p.val {
		return m
	}
	return p.val
}

// ② 求链表的节点个数
func (p *Node) cnt() int {
	if p.Next == nil {
		return 1
	}

	return 1 + p.Next.cnt()
}

// ③ 求所有整数的平均值
func (p *Node) avg(n int) float64 {
	// 当只剩最后一个节点时，本值即是平均值 :从最后一个节点往前看，即归的过程
	if p.Next == nil {
		return float64(p.val)
	}

	avg := p.Next.avg(n - 1)
	return (avg*float64((n-1)) + float64(p.val)) / float64(n)
}

func main() {
	p := Init(1, 3, 2)
	fmt.Println(p.cnt())
	fmt.Println(p.max())
	fmt.Println(p.avg(3))

	fmt.Scanln()
}
