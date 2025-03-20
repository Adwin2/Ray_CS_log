// 难点：分割线的模拟 startIndex
package main

import (
	"fmt"
)

var res [][]string

func backTrack(s string, path []string, startIndex int) {
	// startIndex 脱离了字符串的索引范围 说明分割结束
	if startIndex >= len(s) {
		res = append(res, path)
		return
	}

	// 纵向遍历 更新startIndex
	for i := startIndex; i < len(s); i++ {
		if isPalindrome(s, startIndex, i) {
			// 切片左闭右开
			path = append(path, s[startIndex:i+1])
		} else {
			continue // i ++ 向右寻找
		}
		backTrack(s, path, i+1)   // 寻找 i+1 起始的子串
		path = path[:len(path)-1] // 弹出上一次处理结果 回溯
	}

}

func isPalindrome(s string, start, end int) bool {
	for i, j := start, end; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func combine(startIndex int) {
	path := make([]string, 0)
	backTrack("aab", path, startIndex)
}

func main() {
	combine(0)
	fmt.Println(res)

	//s := "ababa"
	//fmt.Println(isPalindrome(s, 0, len(s)-1))
}
