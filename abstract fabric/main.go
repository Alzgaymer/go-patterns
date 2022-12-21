package main

import "fmt"

func main() {
	s := GetFabric("Adidas")
	a := s.makeShirt("adidasshirt", 30)
	fmt.Printf("Shirt: %v(type: %#t)\n", a, a)
	s = GetFabric("Nike")
	a = s.makeShirt("nikeshirt", 30)
	fmt.Printf("Shirt: %v(type: %#t)\n", a, a)
}
