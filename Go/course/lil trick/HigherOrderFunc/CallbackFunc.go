package main

import (
	"fmt"
	"time"
)

type Callback func()

func asyncOperation(callback Callback) {
	//模拟异步操作
	time.Sleep(2 * time.Second)
	callback()
}

func main() {
	fmt.Println("开始异步操作")
	asyncOperation(func() {
		fmt.Println("异步操作完成")
	})
	fmt.Println("异步操作已启动")
}
