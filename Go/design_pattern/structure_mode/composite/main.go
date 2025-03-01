package main

func main() {
	f1 := &File{name: "file_1"}
	f2 := &File{name: "File_2"}
	f3 := &File{name: "File_3"}

	dir1 := &Folder{name: "dir_1"}
	dir2 := &Folder{name: "dir_2"}

	dir1.AddComposite(f1)
	dir1.AddComposite(f2)
	dir1.AddComposite(f3)

	dir2.AddComposite(f1)
	dir2.AddComposite(f2)
	dir2.AddComposite(f3)

	dir2.Search("content_1")
	dir1.Search("content_2")
}
