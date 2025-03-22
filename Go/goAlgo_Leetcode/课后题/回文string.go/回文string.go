package main

import (
	"fmt"

	"github.com/Adwin2/personal-go-modules/stack"
)

func Verify(a string, s stack.Stack[string]) bool {
	// index
	var i int
	// 现代化写法中类似的： i = range len(a)/2
	// 区别：↑ i 不会执行最后一个++, 以len(a)/2 - 1 返回
	// ↓ i 以len(a)/2 结束返回
	for i = 0; i < len(a)/2; i++ {
		s.Push(string(a[i]))
	}
	//fmt.Println(s.Items)
	//fmt.Println("for循环结束后：", string(a[i]), i)

	if len(a)%2 != 0 {
		//fmt.Println("i++")
		//fmt.Println("之前：", string(a[i]))
		i++
		//fmt.Println("之后：", string(a[i]))
	}

	for ; i < len(a); i++ {
		top, _ := s.Pop()
		//fmt.Println(string(a[i]))
		//fmt.Println(top)
		if string(a[i]) != top {
			return false
		}
	}
	// for i := len(a) - len(a)/2; i < len(s); i++ {
	// 	var ok bool
	// 	if a[i], ok = s.Peek();!ok {
	// 		s.Pop()
	// 	} else {
	// 		return false
	// 	}
	// }
	return s.IsEmpty()
}

func main() {
	s := stack.NewStack[string]()
	fmt.Println(Verify("ababa", *s))
	fmt.Println(Verify("abab", *s))
	fmt.Println(Verify("abba", *s))
	fmt.Println(Verify("ababc", *s))
}
