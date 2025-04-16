package main

import (
	"fmt"
	"os"
)

func main() {
	s := []string{"ray", "ryan"}
	Output1 := fmt.Sprintf("*%d\r\n", len(s))
	for i := 0;i<len(s); i++ {
		Output2 := fmt.Sprintf("$%d\r\n%s\r\n",len(s[i]), s[i])  
		Output1 += Output2
	}
	// ByteOP := []byte(Output1)
	// fmt.Printf("len: %d, ctt: %s", len(ByteOP), ByteOP)
	//str := []byte("*2\r\n$3ray\r\n$4ryan\r\n")
	// var m sync.Map
	// m.Store(2,4)
	// m.Store(3,5)
	// fmt.Println()
	
	os.Create("/home/raymond"+"/"+"a.rdb")

}