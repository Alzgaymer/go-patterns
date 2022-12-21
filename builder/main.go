package main

import "fmt"

func main() {

	s := getBuilder("normal")
	/*set house parameters to (s)*/
	d := newDirector(s)
	d.Build()

	fmt.Printf("%#v", s.GetResult())
}
