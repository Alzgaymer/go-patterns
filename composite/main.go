package main

func main() {
	file1 := File{name: "pisya"}
	file2 := File{name: "popa"}
	file3 := File{name: "kaka"}

	Folder1 := Folder{name: "isporozhneniya"}
	Folder2 := Folder{name: "organi"}

	Folder1.add(&file3)
	Folder2.add(&file1)
	Folder2.add(&file2)

	Folder1.search("popka")
	Folder2.search("piska")
}
