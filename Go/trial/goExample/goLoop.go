package main

import (
	"log"
	"os"
	"sync"
	"thumbnail"
)

func main() {

}

func makeThumbnails5(filenames []string) (thumbfiles []string, err error) {
	type item struct {
		thumbfile string
		err       error
	}

	ch := make(chan item, len(filenames))
	for _, f := range filenames {
		go func(f string) {
			var it item
			//	it.thumbfile, it.err = thumbnail.
		}()
	}
}

/*
	完全彼此独立并行的事件 易并行问题
	未等待线程完成就返回

	可以改变goroutine里的代码让其能够将完成情况报告给外部的goroutine知晓，使用的方式是向一个共享的channel中发送事件。因为我们已经确切地知道有len(filenames)个内部goroutine,所以外部的goroutine只需要在返回之前对这些事件计数。
	使用无缓冲channel时存在error直接返回 发送端channel不会停止 可能导致阻塞或者out of memory （称作goroutine泄漏 
	--解决：因为已知图片（传入文件个数） 所以创建buffered channel


	为了知道最后一个结束的goroutine时间 需要一种特殊的计数器--在多个goroutine操作时做到安全并且减为0前一直等待  -- sync.WaitGroup
	⬇️  ⬇️  ⬇️  
*/

func makeThumbnails6(filenames <-chan string) int64 {
	sizes := make(chan int64)
	var wg sync.WaitGroup //number of working goroutines
	for f := range filenames {
		wg.Add(1)
		//worker
		go func(f string) {
			defer wg.Done()
			thumb, err := thumbnail.ImageFile(f)
			if err != nil {
				log.Println(err)
				return
			}
			info, _ := os.Stat(thumb)
			sizes <- info.Size()
		}(f)
	}
	//closer
	go func() {
		wg.wait()
		close(sizes)
	}()
	var total int64
	for size := range sizes {
		total += size
	}
	return total
}

/*
	- Add为计数器+1 必须在worker routine前调用 不然无法确定add是在closer的wait之前被调用。
	- Done等价于Add(-1)
	- defer确认出错情况也可以正确减去1
	该代码结构是 未知迭代次数的并发循环 的常用写法

	- closer必须在新routine中开  wait贯穿程序 最后close
*/
