package main

import (
	"fmt"
	"slices"
)

var res [][]int

// 树层、树枝去重 used数组
func backtrack(nums, path, used []int) {
	if len(path) == len(nums) {
		res = append(res, slices.Clone(path))
		return
	}

	for i := 0; i < len(nums); i++ {
		//i>0时 used[i-1] == false 同树层使用过 used[i-1] == true 同树枝使用过 -- 树层去重 同理 true时树枝去重 但效率不如树层 (对应0 / 1
		if i > 0 && used[i-1] == 0 && nums[i] == nums[i-1] {
			continue
		}
		if used[i] == 0 {
			used[i] = 1
			path = append(path, nums[i])
			backtrack(nums, path, used)
			used[i] = 0
			path = path[:len(path)-1]
		}
	}
}

func main() {
	nums := []int{1, 1, 2}
	path := make([]int, 0)
	used := make([]int, 10)

	backtrack(nums, path, used)
	fmt.Println(res)
}
