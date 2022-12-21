package main

import "fmt"

type VeggeMania struct {
}

func (v *VeggeMania) getPrice() int {
	const veggemaniaprice = 15
	fmt.Printf("veggemania cost: %v\n", veggemaniaprice)
	return veggemaniaprice
}
