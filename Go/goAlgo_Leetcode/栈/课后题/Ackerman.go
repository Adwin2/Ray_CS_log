package main

import "fmt"

func Ack(m, n int) int {
	if m == 0 {
		return n + 1
	} else if m != 0 && n == 0 {
		return Ack(m-1, 1)
	} else if m != 0 && n != 0 {
		return Ack(m-1, Ack(m, n-1))
	}

	// for nothing
	return 0
}

func main() {
	fmt.Println(Ack(2, 1))
}
