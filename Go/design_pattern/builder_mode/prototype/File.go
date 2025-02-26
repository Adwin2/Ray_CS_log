package main

import "fmt"

type File struct {
	Name string
}

func (f *File) Show(str string) {
	fmt.Println(str + f.Name)
}

func (f *File) Clone() INode {
	return &File{Name: f.Name + "_clone"}
}
