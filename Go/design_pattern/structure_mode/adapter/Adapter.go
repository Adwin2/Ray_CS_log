package main

import "fmt"

type Adapter struct {
	Othertmp *Other
}

func (a *Adapter) InsertUSB() {
	//接口不相容的Linux与Other 可以在这里实现合作
	fmt.Println("Converting ...")
	a.Othertmp.InsertTypeC()
}
