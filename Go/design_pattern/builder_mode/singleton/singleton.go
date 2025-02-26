// 单例模式
// 		懒惰式 ：使用时实例化  once.Do 确保只加载一次
// 		饥饿式 ：初始实例化

package main

import (
	"fmt"
	"sync"
)

var (
	instance *Singleton
	once     = sync.Once{}
)

type Singleton struct{}

func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
	})
	return instance
}

func main() {
	Foo := GetInstance()
	Bar := GetInstance()

	fmt.Printf("Foo = %p\n", Foo)
	fmt.Printf("Bar = %p\n", Bar)
	//Foo = 0x58f380
	//Bar = 0x58f380
}
