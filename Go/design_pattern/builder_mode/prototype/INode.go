package main

type INode interface {
	Show(string)
	Clone() INode
}

// 实现这个接口的都符合 Child []INode
