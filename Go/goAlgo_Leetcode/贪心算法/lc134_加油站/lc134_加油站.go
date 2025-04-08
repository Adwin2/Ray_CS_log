package main

import "fmt"

/*
134. 加油站

题目描述:
在这里添加题目描述

解题思路:
贪心 局部最优推全局最优 且没有反例 -> 贪心

    ① totalSum >= 0 一定可以走完一圈
    ② 出现curSum[i] < 0 起始点更新到 i+1
*/

func solution(gas, cost []int) int {
	start, curSum, totalSum := 0, 0, 0

	// Note : i:= range len(gas) 的索引值是1~len(gas) 并非从0开始
	for i := 0; i < len(gas); i++ {
		curSum += (gas[i] - cost[i])
		totalSum += (gas[i] - cost[i])
		if curSum < 0 {
			start = i + 1
			curSum = 0
		}
	}
	if totalSum < 0 {
		return -1
	}
	return start
}

func main() {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}
	fmt.Println(solution(gas, cost))

	_, _ = fmt.Scanln()
}
