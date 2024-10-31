##死锁 和 goroutine泄漏问题
go func(){
/*
   	实现
   */
close(chan)
}()

告诉接受者goroutine 所有数据都已被全部发送时才需要关闭channel。当chan没有被引用时会被Go语言的垃圾自动回收期回收

##试图重复关闭一个channel将导致panic异常，\试图关闭一个nil值channel也将导致panic异常，\关闭一个channels还会出发一个广播机制。

channel只作为一个函数参数时，一般总是专门用于只发送或者只接收  -- 单方向的channel

chan <- int 只发送int的channel 不能接收
<-chan int 只接收int的channel 不能发送
编译器检测 限制条件.

（对一个只接收的channel调用close将是一个编译错误
