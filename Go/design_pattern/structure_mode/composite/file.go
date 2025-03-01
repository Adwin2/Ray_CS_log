package main

import "fmt"

type File struct {
	name string
}

func (f *File) Search(str string) {
	fmt.Printf("Search file name :[%s], content :[%s]\n", f.name, str)
}

func (f *File) GetName() string {
	return f.name
}
