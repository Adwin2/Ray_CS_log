package main

import (
	"fmt"
	"slices"
)

var res [][]int

// 不回头：startIndex     map[int]struct{} 代unordered_set 避免重复选取
func backtrack(nums, path []int, startIndex int) {
	if len(path) > 1 {
		res = append(res, slices.Clone(path))
	}

	// nums无序 无法使用startIndex来去重，只能用used (unordered_set) 记录已使用元素
	// 只同树层去重时使用 每次进入新树层（递归）(used数组)都会重新初始化
	used := make(map[int]struct{}, 10)
	// for (同树层遍历): 跳过特定元素 +
	for i := startIndex; i < len(nums); i++ {
		_, exists := used[nums[i]]
		if exists || (len(path) > 0 && nums[i] < path[len(path)-1]) {
			continue
		}

		// 数据小：哈希优化：数组used[201]代替map 0/1代替状态
		used[nums[i]] = struct{}{}
		path = append(path, nums[i])
		// 取过的不再遍历 : i+1
		backtrack(nums, path, i+1)
		path = path[:len(path)-1]
	}
}

func main() {
	nums := []int{4, 7, 6, 7}
	path := make([]int, 0)

	backtrack(nums, path, 0)

	fmt.Println(res)
	fmt.Scanln()
}
