package main

import "fmt"

type Strategy func(int, int) int

// 加法策略
func add(a, b int) int {
	return a + b
}

// 减法策略
func substract(a, b int) int {
	return a - b
}

// 使用策略的函数
func executeStrategy(strategy Strategy, a, b int) int {
	return strategy(a, b)
}

func main() {
	result1 := executeStrategy(add, 10, 5)
	fmt.Println("10 + 5 = ", result1)

	result2 := executeStrategy(substract, 10, 5)
	fmt.Println("10 - 5 = ", result2)
}
