package main

import (
	"errors"
	"fmt"

	"github.com/Adwin2/personal-go-modules/stack"
)

type DblStack struct {
	top, bot    [2]int
	left_stack  *stack.Stack[int]
	right_stack *stack.Stack[int]
	m           int
}

// 双栈初始化
// 直接使用  top[0] 作为left_stack.Top
// top[1] 作为 right_stack.Top
func NewDblStack(m int) *DblStack {
	return &DblStack{
		top:         [2]int{-1, m},
		bot:         [2]int{0, m - 1},
		left_stack:  stack.NewStack[int](),
		right_stack: stack.NewStack[int](),
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
func (d *DblStack) Push(elem int, left bool) error {
	if d.isFull() {
		return errors.New("DblStack is Full")
	}

	if left {
		d.left_stack.Push(elem)
		d.left_stack.Top++
		d.top[0]++
	} else {
		d.right_stack.Push(elem)
		d.right_stack.Top++
		d.top[1]--
	}

	return nil
}

// 出栈
func (d *DblStack) Pop(left bool) (int, error) {
	if d.isEmpty() {
		return 0, errors.New("DblStack is Empty")
	}
	var tmp int
	if left {
		tmp, _ = d.left_stack.Peek()
		d.left_stack.Pop()
		d.left_stack.Top--
		d.top[0]--
	} else {
		tmp, _ = d.right_stack.Peek()
		d.right_stack.Pop()
		d.right_stack.Top--
		d.top[1]++
	}

	return tmp, nil
}

// 测试
func main() {
	d := NewDblStack(2)
	// 左插
	fmt.Println((d.top[0]-d.bot[0]+1)+(d.bot[1]-d.top[1]+1), d.m)
	fmt.Println(d.Push(2, true))
	// 右插
	fmt.Println((d.top[0]-d.bot[0]+1)+(d.bot[1]-d.top[1]+1), d.m)
	fmt.Println(d.Push(3, false))

	// 再次插入 已满 报错
	fmt.Println((d.top[0]-d.bot[0]+1)+(d.bot[1]-d.top[1]+1), d.m)
	fmt.Println(d.Push(4, false))

	fmt.Println((d.top[0]-d.bot[0]+1)+(d.bot[1]-d.top[1]+1), d.m)
	fmt.Println(d.Pop(true))
	fmt.Println(d.Pop(false))

	fmt.Println((d.top[0]-d.bot[0]+1)+(d.bot[1]-d.top[1]+1), d.m)
	fmt.Println(d.Push(4, false))
	fmt.Scanln()
}
