package main

import "fmt"

/*
53. 最大子序和

题目描述:
在这里添加题目描述

解题思路:
在这里添加解题思路
*/

func maxSubArray(nums []int) int {
	cnt := 0
	res := 0
	for i := range len(nums) {
		cnt += nums[i]
		// 选起点
		if cnt < 0 {
			cnt = 0
		}

		// 选终点
		if cnt > res {
			res = cnt
		}
	}
	return res
}

func main() {
	// 测试代码
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
}
