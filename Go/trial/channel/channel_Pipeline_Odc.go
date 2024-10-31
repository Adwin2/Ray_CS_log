package main

import "fmt"

func counter(out chan<- int) {
	for x := 0; x < 10; x++ {
		out <- x
	}
	close(out) //关闭发送者
}

func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

func printer(in <-chan int) {
	for z := range in {
		fmt.Println(z)
	}
}

func main() {
	naturals := make(chan int)
	squares := make(chan int)
	go counter(naturals) //隐式类型转换 chan int -- chan<-int
	//双向channel对单向channel变量的赋值都会导致隐式转换   不可从单向到双向
	go squarer(squares, naturals)
	printer(squares)
}
