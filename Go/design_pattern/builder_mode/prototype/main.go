package main

func main() {
	f1 := &File{"文件1"}
	f2 := &File{"文件2"}

	dir1 := &Dir{Name: "目录1", Child: []INode{f2, f2.Clone()}}
	dir2 := &Dir{Name: "目录2", Child: []INode{f1, dir1, f2, dir1.Clone()}}

	dir2.Show("  ")
}
