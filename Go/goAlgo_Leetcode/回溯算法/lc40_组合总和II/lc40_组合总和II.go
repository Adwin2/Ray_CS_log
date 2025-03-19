// candidates有重复、path中同一元素只能使用一遍 : 同树层不可重复选取，同树枝可存在重复选取
// 避免同树层重复选取：① used数组(通用) ② startIndex (i>startIndex: 同树层条件)
package main

import (
	"fmt"
	"slices"
)

var res [][]int

func backTrack(targetSum, sum, startIndex int, candidates, path []int, used []bool) {
	if sum == targetSum {
		res = append(res, slices.Clone(path))
		return
	}

	// 记得剪枝: sum + candidates[i] <= targetSum
	for i := startIndex; i < len(candidates) && sum+candidates[i] <= targetSum; i++ {
		// 跳过同树层使用过的元素  -- for 循环：同树层； 递归调用：延伸树枝
		// ② startIndex方案 跳过条件: i > startIndex && candidates[i] == candidates[i-1]
		if i > 0 && candidates[i] == candidates[i-1] && used[i-1] == false {
			continue
		}

		sum += candidates[i]
		path = append(path, candidates[i])
		used[i] = true
		backTrack(targetSum, sum, i+1, candidates, path, used)
		sum -= candidates[i]
		path = path[:len(path)-1]
		used[i] = false
	}
}

func combine(targetSum, startIndex int) {
	path := make([]int, 0)
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	slices.Sort(candidates)
	used := make([]bool, len(candidates))

	backTrack(targetSum, 0, startIndex, candidates, path, used)
}

func main() {
	combine(8, 0)
	fmt.Println(res)
}
