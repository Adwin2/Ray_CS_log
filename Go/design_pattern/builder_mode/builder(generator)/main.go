package main

import "fmt"

func main() {
	b1 := GetBuild("ice")
	if b1 == nil {
		fmt.Println("not support buildType ...")
		return
	}
	m1 := NewManager(b1)
	iceHouse := m1.GetHouse()
	fmt.Println("iceHouse == ", iceHouse)

	b2 := GetBuild("wood")
	if b2 == nil {
		fmt.Println("not support buildType ...")
		return
	}
	m2 := NewManager(b2)
	woodHouse := m2.GetHouse()
	fmt.Println("woodHouse == ", woodHouse)

	m2.SetWorker(b1)
	Foohouse := m2.GetHouse()
	fmt.Println("Foohouse == ", Foohouse)
}
