package main

import "fmt"

type Cashier struct {
	next IDepartment
}

func (c *Cashier) execute(p *Patient) {
	if p.paymentDone {
		fmt.Println("Payment Done")
	}
	fmt.Println("Cashier getting money from patient patient")
}

func (c *Cashier) setNext(next IDepartment) {
	c.next = next
}
