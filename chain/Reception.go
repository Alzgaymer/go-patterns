package main

import "fmt"

type Reception struct {
	next IDepartment
}

func (r *Reception) execute(p *Patient) {
	if p.registrationDone {
		fmt.Println("Patient registration already done")
		r.next.execute(p)
		return
	}
	fmt.Println("Reception registering patient")
	p.registrationDone = true
	r.next.execute(p)
}
func (r *Reception) setNext(next IDepartment) {
	r.next = next
}
