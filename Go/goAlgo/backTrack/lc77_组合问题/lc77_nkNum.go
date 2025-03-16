package main

import "fmt"

var res [][]int

func backTrack(n, k, StartIndex int, path []int) {
	if len(path) == k {
		NewPath := make([]int, len(path))
		copy(NewPath, path)
		res = append(res, NewPath)
		return
	}
	//fmt.Println(n, k, len(path))
	for i := StartIndex; i <= n-(k-len(path))+1; i++ {
		path = append(path, i)
		backTrack(n, k, i+1, path)
		path = path[:len(path)-1]
	}
}

func combine(n, k int) {
	var nums []int
	backTrack(n, k, 1, nums)
}

func main() {
	combine(4, 2)
	fmt.Println(res)
}
