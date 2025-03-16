package main

import (
	"fmt"
	"sync"
	"time"
)

type SimpleQueue struct {
	queue chan string
	wg    sync.WaitGroup
}

func NewSimpleQueue(capcity int) *SimpleQueue {
	return &SimpleQueue{
		queue: make(chan string, capcity), //带缓存
	}
}

func (s *SimpleQueue) Produce(msg string) {
	s.queue <- msg
	fmt.Println("SimpleQueue Produce: ", msg)
}

func (s *SimpleQueue) Consume(name string) {
	s.wg.Add(1)
	go func() {
		defer s.wg.Done() // select msg :=  <-queue
		for msg := range s.queue {
			time.Sleep(500 * time.Millisecond) //模拟处理时间
			fmt.Printf("[消费者%s]完成消息接收：%s \n", name, msg)
		}
	}()
}

func (s *SimpleQueue) Close() {
	close(s.queue)
	s.wg.Wait()
}

func main() {
	s := NewSimpleQueue(2)
	s.Consume("a")
	s.Consume("b")

	go func() {
		for i := 1; i <= 10; {
			msg := fmt.Sprintf("msg %d", i)
			s.Produce(msg)
			time.Sleep(100 * time.Millisecond)
			i++
		}
	}()
	time.Sleep(3 * time.Second)
	s.Close()
	fmt.Println("all done")
}
