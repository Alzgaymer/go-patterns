package main

import "fmt"

type TomatoToping struct {
	Pizza IPizza
}

func (t *TomatoToping) getPrice() int {
	const tomatoTopingPrice = 7
	fmt.Printf("Tomato toping cost: %v\n", tomatoTopingPrice)
	return t.Pizza.getPrice() + tomatoTopingPrice
}
