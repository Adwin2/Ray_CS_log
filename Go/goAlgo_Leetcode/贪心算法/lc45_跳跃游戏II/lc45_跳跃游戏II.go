package main

import "fmt"

/*
45. 跳跃游戏II

题目描述:
在这里添加题目描述

解题思路:
在这里添加解题思路
*/

func minStep(nums []int) int {
	step := 0
	curIdx := 0
	for i := range len(nums) - 1 {
		step += 1
		curIdx = max(i+nums[i], curIdx)
		if curIdx >= len(nums)-1 {
			return step
		}
	}
	return step + 1
}

func main() {
	nums := []int{2, 3, 1, 1, 4}
	fmt.Println(minStep(nums))
}
