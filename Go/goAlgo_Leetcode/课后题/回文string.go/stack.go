package main

// 定义一个泛型的Stack结构体
type Stack[T any] struct {
	Items []T
	Top   int
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		Items: make([]T, 0),
		Top:   -1,
	}
}

// 实现Push方法来添加元素
func (s *Stack[T]) Push(item T) {
	s.Items = append(s.Items, item)
	s.Top++
}

// 实现Pop方法来移除并返回栈顶元素
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.Items) == 0 {
		var zero T // 如果栈为空，返回类型的零值(初始化默认零值)和false
		return zero, false
	}
	item := s.Items[len(s.Items)-1]
	s.Items = s.Items[:len(s.Items)-1]
	s.Top--
	return item, true
}

// 实现Peek方法来查看栈顶元素但不移除它
func (s *Stack[T]) Peek() (T, bool) {
	if len(s.Items) == 0 {
		var zero T // 如果栈为空，返回类型的零值和false
		return zero, false
	}
	return s.Items[len(s.Items)-1], true
}

// 实现IsEmpty方法来检查栈是否为空
func (s *Stack[T]) IsEmpty() bool {
	return len(s.Items) == 0
}

// func main() {
// 	// 使用int类型创建栈实例
// 	stack := Stack[int]{}
// 	stack.Push(1)
// 	stack.Push(2)
// 	stack.Push(3)
// 	fmt.Println(stack.Pop())  // 输出: 3 true
// 	fmt.Println(stack.Peek()) // 输出: 2 true
// 	fmt.Println(stack.IsEmpty()) // 输出: false

// 	// 使用string类型创建栈实例
// 	stackStr := Stack[string]{}
// 	stackStr.Push("Hello")
// 	stackStr.Push("World")
// 	fmt.Println(stackStr.Pop())  // 输出: World true
// 	fmt.Println(stackStr.Peek()) // 输出: Hello true
// 	fmt.Println(stackStr.IsEmpty()) // 输出: false
// }

// --------- v1 不支持泛型 -------------
// package stack

// import (
// 	"errors"
// )

// type Stack struct {
// 	Elems []int
// 	Top   int
// }

// func NewStack(size int) *Stack {
// 	return &Stack{
// 		Elems: make([]int, size),
// 		Top:   -1,
// 	}
// }

// // 入栈
// func (s *Stack) Push(Elem int) error {
// 	if s.isFull() {
// 		return errors.New("Stack is Full")
// 	}
// 	//s.Elems = append(s.Elems, Elem)
// 	s.Top++
// 	s.Elems[s.Top] = Elem
// 	return nil
// }

// // 出栈
// func (s *Stack) Pop() (int, error) {
// 	if s.isEmpty() {
// 		return 0, errors.New("Stack is empty")
// 	}

// 	Elem := s.Elems[s.Top]
// 	s.Elems = s.Elems[:len(s.Elems)-1]
// 	s.Top--

// 	return Elem, nil
// }

// // 查看栈顶元素但不弹出
// func (s *Stack) Peek() int {
// 	if len(s.Elems) == 0 {
// 		return 0
// 	}
// 	return s.Elems[len(s.Elems)-1]
// }

// // 返回栈的大小
// func (s *Stack) Size() int {
// 	return s.Top + 1
// }

// // 判断栈是否为空
// func (s *Stack) isEmpty() bool {
// 	return s.Top < 0
// }

// // 判断栈是否已满
// func (s *Stack) isFull() bool {
// 	return s.Top >= cap(s.Elems)-1
// }
