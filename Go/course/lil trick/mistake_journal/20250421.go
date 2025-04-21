package main

import "fmt"

func main() {
	//	i := 1
	// j := i++
	// 	str := "huhi"
	fmt.Println(trial(1, 2))
}

func trial(a, b int) (c int) {
	// 这里返回值已经变成 a+b 
	defer func() {
		c++
	}()
	// 这里不影响
	c++
	return a + b
}

/*
	1. 自增自减是语句 不是表达式  不可用来赋值
	2. 字符串中的字符不支持取地址
	3. defer函数 对返回值的影响
*/
