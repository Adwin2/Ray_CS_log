package main

import "fmt"

type Linux struct{}

func (m *Linux) InsertUSB() {
	fmt.Println("Linux insert USB ...")
}
