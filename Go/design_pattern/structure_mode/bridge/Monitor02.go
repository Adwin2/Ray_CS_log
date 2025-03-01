package main

import "fmt"

type Monitor02 struct{}

func (m *Monitor02) Show(content string) {
	fmt.Printf("[%s]M02 show ...\n", content)
}
