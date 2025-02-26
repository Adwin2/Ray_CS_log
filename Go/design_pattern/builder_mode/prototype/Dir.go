package main

import "fmt"

type Dir struct {
	Name  string
	Child []INode
}

func (d *Dir) Show(str string) {
	fmt.Println(str + d.Name)
	for _, ch := range d.Child {
		ch.Show(str + str)
	}
}

func (d *Dir) Clone() INode {
	cloneChild := make([]INode, len(d.Child))

	for index, ch := range d.Child {
		cloneChild[index] = ch.Clone()
	}
	return &Dir{Name: d.Name + "_clone", Child: cloneChild}
}
