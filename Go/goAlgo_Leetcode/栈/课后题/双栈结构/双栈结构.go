package main

import "../../stack"

type DblStack struct {
	top, bot    [2]int
	left_stack  *stack.Stack
	right_stack *stack.Stack
	m           int
}

//重载栈的初始化
// func NewStack(size, top int) *stack.Stack{
// 	return &stack.Stack{
// 		Elems: make([]int, size),
// 		Top: top,
// 	}
// }

// 双栈初始化
// 直接使用  top[0] 作为left_stack.Top
// top[1] 作为 right_stack.Top
func NewDblStack(m int) *DblStack {
	return &DblStack{
		top:         [2]int{-1, m},
		bot:         [2]int{0, m - 1},
		left_stack:  stack.NewStack(m),
		right_stack: stack.NewStack(m),
		m:           m,
	}
}

// 判断栈空
func (d *DblStack) isEmpty() bool {
	return d.top == [2]int{-1, d.m}
}

// 判断栈满
func (d *DblStack) isFull() bool {
	return (d.top[0]-d.bot[0]+1)+(d.bot[1]-d.top[1]+1) == d.m
}

// 进栈
func (d *DblStack) Push(elem int) {

}

// 出栈

// 测试
