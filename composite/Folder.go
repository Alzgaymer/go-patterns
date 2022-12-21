package main

import "fmt"

type Folder struct {
	components []IComponent
	name       string
}

func (f *Folder) search(keyword string) {
	fmt.Printf("Searching in folder %s, for \"%s\"\n", f.name, keyword)
	for _, v := range f.components {
		v.search(keyword)
	}
}
func (f *Folder) add(newComponent IComponent) {
	f.components = append(f.components, newComponent)
}
