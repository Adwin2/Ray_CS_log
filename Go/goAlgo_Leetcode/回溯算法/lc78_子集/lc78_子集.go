// 子集中没有重复元素
package main

import (
	"fmt"
	"slices"
)

var res [][]int

func backtrack(path, nums []int, startIndex int) {
	// 一： slices.Clone(Path) 副本入栈 “归”的时候才是正确的值
	res = append(res, slices.Clone(path))
	for i := startIndex; i < len(nums); i++ {
		// 二：i 是索引， append nums[i]
		path = append(path, nums[i])
		backtrack(path, nums, i+1)
		path = path[:len(path)-1]
	}

}

func combine(nums []int) {
	path := make([]int, 0)
	backtrack(path, nums, 0)
}

func main() {
	nums := []int{1, 2, 3}
	combine(nums)
	fmt.Println(res)
}
