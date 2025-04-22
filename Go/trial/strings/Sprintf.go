package main

import (
	"bytes"
	"fmt"
	"sync"
	"time"
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
	
	// os.Create("/home/raymond"+"/"+"a.rdb")
	// cnt := 10
	// fmt.Println(cnt, byte(cnt))
	var key any = "key"
	var t int64 = 200
	fmt.Println(t, time.Now().UnixMilli(),t+time.Now().UnixMilli(), uint64(t)+ uint64(time.Now().UnixMilli()))
	str, _ := key.(string)
	fmt.Println(byte(len(key.(string)))==byte(0x03), []byte(str), bytes.Join([][]byte {
		[]byte{byte(len(key.(string)))},
		[]byte(key.(string)),
	}, nil))
	cnt := 0
	var m sync.Map
	m.Store(1,2)
	m.Store(2,2)
	m.Range(func (k, v any) bool {
		cnt ++
		return true
	})
	fmt.Println(cnt)
	sss := "A"
	fmt.Println(HandleString("adddndnsjs" + sss))
	//panic("cnt < 5")

}

func HandleString(s string) []byte {
	return []byte(s)
}