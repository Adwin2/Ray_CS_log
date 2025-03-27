package main

import (
	"fmt"
	"sort"
)

/*
455. 分发饼干

题目描述:
在这里添加题目描述

解题思路:
小尺寸满足小胃口 逐个比较 (大的同理)
*/

func CntContentChildren(req, size []int) int {
	index := 0
	for i := 0; i < len(size); i++ {
		if index < len(req) && size[i] >= req[index] {
			index++
		}
	}

	return index
}

func main() {
	req := []int{1, 2}
	size := []int{1, 2, 3}
	// sort.Slice(req, func(i, j int) bool {
	// 	return i > j
	// })
	sort.Ints(req)
	sort.Ints(size)
	fmt.Println(CntContentChildren(req, size))
}
