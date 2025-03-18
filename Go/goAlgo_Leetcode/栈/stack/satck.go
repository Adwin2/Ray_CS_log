package stack

import (
	"errors"
)

type Stack struct {
	Elems []int
	Top   int
}

func NewStack(size int) *Stack {
	return &Stack{
		Elems: make([]int, size),
		Top:   -1,
	}
}

// 入栈
func (s *Stack) Push(Elem int) error {
	if s.isFull() {
		return errors.New("Stack is Full")
	}
	//s.Elems = append(s.Elems, Elem)
	s.Top++
	s.Elems[s.Top] = Elem
	return nil
}

// 出栈
func (s *Stack) Pop() (int, error) {
	if s.isEmpty() {
		return 0, errors.New("Stack is empty")
	}

	Elem := s.Elems[s.Top]
	s.Elems = s.Elems[:len(s.Elems)-1]
	s.Top--

	return Elem, nil
}

// 查看栈顶元素但不弹出
func (s *Stack) Peek() (int, error) {
	if len(s.Elems) == 0 {
		return 0, errors.New("Stack is empty")
	}
	return s.Elems[len(s.Elems)-1], nil
}

// 返回栈的大小
func (s *Stack) Size() int {
	return s.Top + 1
}

// 判断栈是否为空
func (s *Stack) isEmpty() bool {
	return s.Top < 0
}

// 判断栈是否已满
func (s *Stack) isFull() bool {
	return s.Top >= cap(s.Elems)-1
}
