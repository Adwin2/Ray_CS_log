/*
并发 多线程在一个CPU运行
并行 多个程序在多个核的CPU运行

线程：内核态 （线程跑多个协程   栈KB级别 
协程 ：用户级的（轻量的线程   栈MB级别

协程之间的通信  go：通过通信来共享内存 －－ channel

并发安全Lock  sync.Mutex  lock() unlock() 

sync.WaitGroup  Add()开启协程 \ Done()执行结束 \ Wait() 主协程阻塞至计数器为0 （所有协程运行结束）
*/
package main

func main() {

}
