package main

import "fmt"

func main() {
	pizza := &VeggeMania{}
	pizzaWithTomato := &TomatoToping{Pizza: pizza}
	fmt.Printf("Pizza with tomato: %v\n", pizzaWithTomato.getPrice())
	fmt.Println("____________________________________________________")
	pizzaWithTomatoAndCheese := &CheeseToping{Pizza: pizzaWithTomato}
	fmt.Printf("Pizza with tomato and cheese: %v\n", pizzaWithTomatoAndCheese.getPrice())
	fmt.Println("____________________________________________________")

}
