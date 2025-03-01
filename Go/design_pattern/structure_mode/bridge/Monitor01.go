package main

import "fmt"

type Monitor01 struct{}

func (m *Monitor01) Show(content string) {
	fmt.Printf("[%s]M01 show ...\n", content)
}
