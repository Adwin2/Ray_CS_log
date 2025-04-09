package main

import "fmt"

/*
135. 分发糖果

题目描述:
在这里添加题目描述

解题思路:
在这里添加解题思路
*/

func solution(ratings []int) int {
	candyVec := make([]int, len(ratings))
	for i := 0; i < len(candyVec); i++ {
		candyVec[i] = 1
	}

	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			candyVec[i] = candyVec[i-1] + 1
		}
	}
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			candyVec[i] = max(candyVec[i], candyVec[i+1]+1)
		}
	}

	res := 0
	for i := 0; i < len(candyVec); i++ {
		res += candyVec[i]
	}
	return res
}

func main() {
	test1 := []int{1, 0, 2}
	test2 := []int{1, 2, 2}

	fmt.Println(solution(test1))
	fmt.Println(solution(test2))
}
