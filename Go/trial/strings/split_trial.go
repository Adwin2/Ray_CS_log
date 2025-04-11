package main

import (
	"fmt"
	"log"

	// "github.com/schollz/progressbar/v3"
	"strconv"
	"strings"
	"time"
)

func main() {
	s := "*2\r\n$4\r\nECHO\r\n$3\r\nhey\r\n"

	res := strings.Split(s, "\r\n")
	if strings.HasPrefix(res[0], "*") {
		elem_Len, err := strconv.Atoi(res[0][1:])
		if err != nil {
			return
		}
		if strings.EqualFold(res[1*elem_Len], "echo") {
			log.Printf("Received ECHO, Output:%s", res[2*elem_Len])
			// conn.Write("+%s\r\n", res[2*elem_Len])
		}
	} else {
		log.Printf("Received PING")
		// conn.Write("+PONG\r\n")
	}
	fmt.Println(res[:len(res)-1], len(res))
	t := 300
	if len(res) > 6 && res[100] == "0" {
		log.Printf("test")
	}
	time.Sleep(time.Duration(t) * time.Millisecond)
	log.Printf("waiting...")
}

func handleECHO() {

}
