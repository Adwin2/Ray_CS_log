package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	worklist := make(chan []string)

	//start with the command-line arguments
	//  var n int\ n++   ---for crawl2 有传入参数就加一
	go func() {
		worklist <- os.Args[1:]
	}()

	//crawl web concurrently
	seen := make(map[string]bool)
	//	for;n > 0;n--   ---for crawl2
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				//n++ ---for crawl2 每启动crawler之前
				go func(link string) {
					worklist <- crawl(link)
				}(link) //显式传入link参数 避免循环变量快照问题 - 读取命令行参数和crawler 分为两个routine（都是发送routine 避免死锁
			}
		}
	}
}

func crawl1(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

/*
⬆️  ⬆️  ⬆️
	crawl1问题：一次性创建了太多网络连接 ，超过了进程打开文件数限制（系统限制因素）
	--通过限制程序使用资源 来适应运行环境  
	--〉避免过度并发〈--
*/

// 思路一 计数信号量
var token = make(chan struct{}, 20) //chan 内类型不重要 -struct{} 20：inode限制

func crawl2(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{}
	list, err := links.Extract(url)
	<-tokens //release the token
	if err != nil {
		log.Print(err)
	}
	return list
}

// 思路二 常驻多个routine
func main() {
	worklist := make(chan []string)
	unseenLinks := make(chan string)

	go func() {
		worklist <- os.Args[1:]
	}()

	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() {
					worklist <- foundLinks //抓完发送到 worklist避免死锁 why new goroutine: unseenLinks 和 worklist都是满的情况下 交叉等待造成死锁 接收与抓取分开执行 提高效率
				}()
			}
		}()
	}

	//the main routine de-duplicates worklist items and send the unseen ones to the crawlers.
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		} //把没有抓过的发给crawler
	}

	//map被限制在main routine 中  -- 内部变量无法访问到 （信息隐藏）
	// 〈拓〉： 变量逃逸-- 局部变量被全局变量引用地址导致变量被分配到堆上
	//	一个对象的封装字段无法在该对象方法外访问
}
