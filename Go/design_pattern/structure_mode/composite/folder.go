package main

import "fmt"

type Folder struct {
	composite []Composite
	name      string
}

func (f *Folder) Search(content string) {
	fmt.Printf("searching folder :[%s],content :[%s]\n", f.name, content)
	for _, v := range f.composite {
		v.Search(content)
	}
}

func (f *Folder) AddComposite(c Composite) {
	f.composite = append(f.composite, c)
}
