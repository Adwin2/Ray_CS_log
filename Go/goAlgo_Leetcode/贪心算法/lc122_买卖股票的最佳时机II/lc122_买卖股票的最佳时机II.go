package main

import "fmt"

/*
122. 买卖股票的最佳时机II

题目描述:
在这里添加题目描述

解题思路:
在这里添加解题思路
*/
func maxProfit(nums []int) int {
	cur := 0
	res := 0
	for i := 1; i < len(nums); i++ {
		cur = nums[i] - nums[i-1]
		if cur > 0 {
			res += cur
		}
	}

	return res
}

func main() {
	// 测试代码
	nums := []int{7, 1, 5, 3, 6, 4}
	//nums := []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit(nums))
}
