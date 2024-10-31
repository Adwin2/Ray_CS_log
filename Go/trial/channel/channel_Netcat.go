package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn) //Note:ignoring errors
		log.Println("done")
		done <- struct{}{} //signal the main goroutine 可以使用bool或int类型实现同样的功能 a send statement (to main goroutine(where its created)
	}()
	mustCopy(conn, os.Stdin)
	conn.Close()
	<-done //wait for background goroutine to finish (discard the receive result)
}
