package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		}
	}
} //channel buffer为1 select的两个case只有一个执行 （一个接收一个发送 满和空的状态交替出现）

/*
多channel接收信息 阻塞-- select multiplex（多路复用）
	如果多个case同时就绪，select会随机选择一个执行
	空select{}会一直等待
*/
