package main

import (
	"fmt"
	"sync"
	"time"
)

type Callback func()

var wg sync.WaitGroup

func asyncOperation(callback Callback) {
	//模拟异步操作
	time.Sleep(2 * time.Second)
	callback()
	wg.Done()
}

func main() {
	fmt.Println("开始异步操作")

	wg.Add(1)
	go asyncOperation(func() {
		fmt.Println("异步操作完成")
	})
	fmt.Println("异步操作已启动")
	wg.Wait()
}
