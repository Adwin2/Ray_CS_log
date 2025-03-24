package main

import (
	"fmt"
	"strings"
)

var res []string

// startIndex : 每字符仅遍历一遍(i+1)   pointNum 结束条件
func backtracks(s string, startIndex, pointNum int) {
	if pointNum == 3 {
		//fmt.Println("final step：", s)
		if isValid(s, startIndex, len(s)-1) {
			res = append(res, s)
		}
		// 直接返回
		return
	}

	for i := startIndex; i < len(s); i++ {
		if isValid(s, startIndex, i) {
			pointNum++
			// insert(s, i+1)
			s = strings.Join([]string{s[:i+1], ".", s[i+1:]}, "")
			//fmt.Println("回溯前 s:", s)
			backtracks(s, i+2, pointNum)
			pointNum--
			s = s[:i+1] + s[i+2:]
			//fmt.Println("回溯 s:", s)
		} else {
			// 出现不合法 本层循环直接结束
			break
		}
	}
}

func isValid(s string, start, end int) bool {
	if start > end {
		return false
	}
	// 段首是'0'
	if s[start] == '0' && start != end {
		return false
	}

	var num int64
	for i := start; i <= end; i++ {
		// 非正整数
		if s[i] > '9' && s[i] < '0' {
			return false
		}
		num = (num*10 + int64(s[i]-'0'))
	}
	// 段总和大于255
	if num > 255 {
		return false
	}

	return true
}

func combine(s string) {
	backtracks(s, 0, 0)
}

func main() {
	combine("025511135")

	fmt.Println(res)
	fmt.Scanln()
}
