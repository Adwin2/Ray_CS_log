// 子集中有重复元素
package main

import (
	"fmt"
	"slices"
)

var res [][]int

// 使用startIndex 集合 不回头取元素
func backtrack(nums, path []int, startIndex int, used map[int]bool) {
	res = append(res, slices.Clone(path))
	for i := startIndex; i < len(nums); i++ {
		// 同树层去重  前一个相同元素已回溯过
		// used 数组方法较通用 如排列问题 此处i > startIndex 也可以替代条件一三
		if i > 0 && nums[i] == nums[i-1] && used[i-1] == false {
			continue
		}
		if used[i] == false {
			used[i] = true
			path = append(path, nums[i])
			backtrack(nums, path, i+1, used)
			used[i] = false
			path = path[:len(path)-1]
		}
	}
}

func main() {
	nums := []int{1, 2, 2}
	path := make([]int, 0)
	used := make(map[int]bool, len(nums))
	backtrack(nums, path, 0, used)

	fmt.Println(res)
}
