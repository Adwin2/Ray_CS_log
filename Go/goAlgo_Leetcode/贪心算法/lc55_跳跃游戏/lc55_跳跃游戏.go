package main

import "fmt"

/*
55. 跳跃游戏

题目描述:
给定一个非负整数数组，你最初位于数组的第一个位置。
数组中的每个元素代表你在该位置可以跳跃的最大长度。
判断你是否能够到达最后一个位置。

解题思路:
使用贪心算法，维护一个变量 cover 表示当前能够到达的最远位置。
遍历数组，更新 cover 的值为 max(cover, i + nums[i])。
如果在某次更新后 cover 已经大于或等于最后一个位置的下标，返回 true。
如果遍历结束后 cover 仍然小于最后一个位置的下标，返回 false。
*/

// 移动的时候更新cover (idx)能 >= 终点idx就可以到达
func canGetEnd(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	// 相当于下标
	cover := 0
	// 遍历到倒数第二个元素
	for i := 0; i < len(nums)-1; i++ {
		cover = max(i+nums[i], cover)
		if cover >= len(nums)-1 {
			return true
		}
	}
	return false
}

func main() {
	nums := []int{2, 3, 1, 1, 4}
	nums0 := []int{3, 2, 1, 0, 4}
	fmt.Println(canGetEnd(nums))
	fmt.Println(canGetEnd(nums0))
}
