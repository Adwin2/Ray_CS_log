package main

import "fmt"

func test_a(s string) int {
	num := 0
	for i := range len(s) {
		fmt.Println(int(s[i] - '0'))
		num = (num*10 + int(s[i]-'0'))
		fmt.Println("num:", num)
	}
	return num
}

func main() {
	s := "12345"
	// n := s[1] - '0'
	// fmt.Println(n)
	// fmt.Println(s[:1] + s[2:])
	fmt.Println(test_a(s))
}
