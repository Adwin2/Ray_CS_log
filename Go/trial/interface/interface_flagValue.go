package main

import (
	"flag"
	"fmt"
	"time"
)

var period = flag.Duration("period", 1*time.Second, "sleep period")

/*
	flag.Duration 创建一个 time.Duration类型的标记变量
	通过 -period 命令来控制 ./<name> -period 2m30s 
	same with 50ms 1.5h
	（时间周期标记值 该特性被构建到了flag包中
*/

func main() {
	flag.Parse()
	fmt.Printf("Sleeping for %v...", *period)
	time.Sleep(*period)
	fmt.Println()
}
