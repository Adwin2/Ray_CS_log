package main

import "fmt"

func main() {
	f(3)
}

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x) //panic if x == 0
	defer fmt.Printf("defer %d \n", x)
	f(x - 1)
}

//拓展:延迟函数的调用在释放堆栈信息之前
