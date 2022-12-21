package main

import "fmt"

type File struct {
	name string
}

func (f *File) search(keyword string) {
	fmt.Printf("Searching in file %s, for \"%s\"\n", f.name, keyword)
}

func (f *File) getName() string {
	return f.name
}
