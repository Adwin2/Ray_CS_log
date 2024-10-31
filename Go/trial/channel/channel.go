package main

func main() {
	ch := make(chan int) //ch has type "chan int"
	/*零值也是nil 可以比较 ；有发送和接收两个功能 都使用 <- 符号
	ch <- x //a send statement
	x = <- ch //a receive statement
	<- ch //a receive statement; result is discarded
	/使用内置close函数即可关闭
	close(ch)
	/调用make函数创建的是一个无缓存的channel，可以指定第二个整型参数， 对应channel的容量（也就是带缓存的channel）
	ch = make(chan int) //unbuffered channel
	ch = make(chan int, 0)//unbuffered channel
	ch = make(chan int, 3) //buffered channel with capacity 3
	*/

	//无缓存channel 
	/*
		执行发送操作将导致发送者goroutine堵塞，直到另一个goroutine在相同的channel执行接受操作；反之同理
		即导致两个goroutine做一次同步操作
			也叫做同步channels (happens before
	*/
	//单方向的channel

//带缓存的channel***
	/*
			元素队列 发送：向队列尾部插入元素 接收：队列头部弹出（删除）元素
			若内部缓存队列满了要阻塞到下一个接受操作 空队列 同理 接受操作将阻塞 直到另一个goroutine的发送操作插入元素
		如果一个goroutine对一个满的缓存队列接收了一个值 那么该缓存队列下一个操作不会发生阻塞  称 channel的缓存队列解耦了接收和发送的goroutine

		cap(chan) 获取内部缓存容量
		len(chan) 获取有效元素个数

		Channel和goroutine的调度器机制是紧密相连的，如果没有其他goroutine从channel接收，发送者——或许是整个程序——将会面临永远阻塞的风险。如果你只是需要一个简单的队列，使用slice就可以了。

	*/

	//goroutine泄漏问题 --_mirroredQuery()
	/*
		如果使用了无缓存channel 那么两个慢的goroutines将会永远卡住
		称为goroutine泄漏 （不再需要的goroutine不再会被自动回收 要确保正常退出
	*/

无缓存  -- 保证同步操作
带缓存  -- 操作解耦  （大小设定不当易导致死锁  也可能影响程序性能
  //生产线模型
}





/*
func mirroredQuery() string{
	responses:= make(chan string, 3)
	go func() {
		response <- request("asia.gopl.io")
	}()
	go func{
		response <- request("europe.gopl.io")
	}()

	go func() {
		responses <- request("americas.gopl.io")
	}()
	return <-responses //return the quickest response
}

func request(hostname string) (response string) {}
	
*/
