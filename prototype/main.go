package main

import "fmt"

func main() {
	a := Reactengle{X: 2, Y: 3, Color: "red"}

	b := a.clone()
	fmt.Printf("%#v\n", a)
	fmt.Printf("%#v\n", b)

	s := Circle{X: 1, Y: 1, D: 2, Color: "blue"}
	b = s.clone()

	fmt.Printf("%#v\n", a)
	fmt.Printf("%#v\n", b)
	fmt.Printf("%#v\n", s)

}
