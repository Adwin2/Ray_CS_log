package main

import (
	"fmt"
	"slices"
)

var res [][]int

func GroupSum(cur, sum, start, k int, path []int) {
	if cur > sum {
		return
	}
	if len(path) == k {
		if cur == sum {
			//res = append(res, append([]int{}, path...))
			res = append(res, slices.Clone(path))
			return
		}
	}

	for i := start; i <= 9-(k-len(path))+1; i++ {
		cur += i
		path = append(path, i)
		GroupSum(cur, sum, i+1, k, path)
		cur -= i
		path = path[:len(path)-1]
	}
}

func combine() {
	// 	cur := 0
	// 	res = make([][]int, 0)
	path := make([]int, 0)
	// nums := []int{1, 2, 3, 4, 5, 6, 7}
	GroupSum(0, 4, 1, 2, path)
}

func main() {
	combine()
	fmt.Println(res)
}
