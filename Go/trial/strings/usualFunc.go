package main

import (
	"bufio"
	"fmt"
	"os"
)

func Join() string {
	fmt.Scanln()
}
func test() {

	// 默认最多读取65536 bytes 字符 通过scanner.Buffer() 调大缓冲区最大读取大小
	bufio.NewScanner(os.Stdin)
	// Split()  Atoi()
}
