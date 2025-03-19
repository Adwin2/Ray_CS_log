// candidates 无重复（同树层不存在重复）、 path可出现重复 同树枝可重复选取 不限制层数
// 确定回溯参数 、终止条件 、单层搜索逻辑
package main

import (
	"fmt"
	"slices"
)

var res [][]int

func backTrack(targetSum, sum, startIndex int, candidates, path []int) {
	if sum == targetSum {
		res = append(res, slices.Clone(path))
		return
	}

	for i := startIndex; i < len(candidates) && sum+candidates[i] <= targetSum; i++ {
		sum += candidates[i]
		path = append(path, candidates[i])
		backTrack(targetSum, sum, i, candidates, path)
		sum -= candidates[i]
		path = path[:len(path)-1]
	}
}

func combine(targetSum, startIndex int) {
	candidates := []int{2, 3, 6, 7}
	path := make([]int, 0)
	backTrack(targetSum, 0, startIndex, candidates, path)
}

func main() {
	combine(7, 0)
	fmt.Println(res)
	//fmt.Scanln()
}
