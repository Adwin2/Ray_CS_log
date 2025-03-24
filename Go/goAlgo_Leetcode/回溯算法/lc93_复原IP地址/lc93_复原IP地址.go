package main

import "fmt"

var res [][]string

// startIndex : 每字符仅遍历一遍(i+1)   pointNum 结束条件
func backtracks(s string, startIndex, pointNum int) {

}

func combine(s string) {
	backtracks(s, 0, 0)
}

func main() {
	combine("25525511135")

	fmt.Println(res)
	fmt.Scanln()
}
