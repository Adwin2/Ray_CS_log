package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "2\r\ne\r\n"

	res := strings.Split(s, "\r\n")
	fmt.Println(res[:len(res)-1], len(res))
}
