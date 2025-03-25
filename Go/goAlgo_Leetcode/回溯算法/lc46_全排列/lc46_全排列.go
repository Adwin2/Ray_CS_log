// nums不包含重复元素
package main

import (
	"fmt"
	"slices"
)

var res [][]int

// 排列问题不需要startIndex 每树层都是从index=0开始
// 去重使用used数组 : 标记同树枝nums数组已经使用哪些
func backtrack(nums, used, path []int) {
	if len(path) == len(nums) {
		res = append(res, slices.Clone(path))
		return
	}
	for i := 0; i < len(nums); i++ {
		if used[i] == 1 {
			continue
		}
		used[i] = 1
		path = append(path, nums[i])
		backtrack(nums, used, path)
		used[i] = 0
		path = path[:len(path)-1]
	}
}

func main() {
	nums := []int{1, 2, 3}
	path := make([]int, 0)
	used := make([]int, 10)
	backtrack(nums, used, path)

	fmt.Println(res)
	fmt.Scanln()
}
