package main

import "fmt"

type CheeseToping struct {
	Pizza IPizza
}

func (c *CheeseToping) getPrice() int {

	const CheeseTopingPrice = 10
	fmt.Printf("Cheese toping cost: %v\n", CheeseTopingPrice)
	return c.Pizza.getPrice() + CheeseTopingPrice
}
