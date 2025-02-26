package main

import "fmt"

type FilterFunc func(int) bool

type MapFunc func(int) int

// 过滤函数
func filter(numbers []int, f FilterFunc) []int {
	var result []int
	for _, num := range numbers {
		if f(num) {
			result = append(result, num)
		}
	}
	return result
}

// 映射函数
func mapFunc(numbers []int, f MapFunc) []int {
	var result []int
	for _, num := range numbers {
		result = append(result, f(num))
	}
	return result
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//过滤出偶数
	evenNumbers := filter(numbers, func(n int) bool {
		return n%2 == 0
	})
	fmt.Println(evenNumbers)
	//将每个数平方
	squaredNumbers := mapFunc(evenNumbers, func(n int) int {
		return n * n
	})
	fmt.Println(squaredNumbers)
}
