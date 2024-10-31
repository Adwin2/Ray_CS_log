/*
	main 本身就是协程之一 自身的无缓冲channel会使main协程 阻塞
	写入的无缓冲channel必须等到 读的channel才可以继续运行
*/

package main

import "fmt"

func main() {
	/*
		channels可以用于将多个goroutine连接在一起，一个channel的输出可以作为下一个channel的输入 串联的channel即所谓的管道
	*/

	naturals := make(chan int)
	squares := make(chan int)

	//Counter
	go func() {
		for x := 0; x < 10; x++ {
			naturals <- x
		}
		close(naturals)
	}() //并发 传递多个数据会导致死锁 >通过关闭channel来解决 一个被关闭的channel已发送的数据都被成功接受以后 后续的接受操作将不再阻塞，立即返回零值 >复杂程序使用defer(close(chan) + go func(){} () 来实现

	//Squarer
	go func() {
		for {
			x, ok := <-naturals
			if !ok {
				break
			}
			squares <- x * x
		}
		close(squares)
	}()

	//Printer (in main goroutine) range循环可以直接在channel上迭代
	for x := range squares {
		fmt.Println(x)
	}
}
