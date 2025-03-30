package main

import "fmt"

/*
376. 摆动序列

题目描述:
在这里添加题目描述

解题思路:
在这里添加解题思路
*/
// 返回折线的折点数
func solution(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	// 第一个点
	res := 1
	PreDiff, CurDiff := 0, 0
	for i := 1; i < len(nums); i++ {
		CurDiff = nums[i] - nums[i-1]
		// =0第二个点
		if CurDiff*PreDiff <= 0 {
			res++
			//fmt.Printf("%d\t", nums[i])
		}
		PreDiff = CurDiff
	}
	return res
}

func main() {
	// 测试代码
	nums := []int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8}
	nums2 := []int{1, 7, 4, 9, 2, 5}
	fmt.Println(solution(nums))
	fmt.Println(solution(nums2))
}
