package main

import "fmt"

type FreightTrain struct {
	mediator IMediator
}

func (p *FreightTrain) arrive() {
	if p.mediator.canArrive(p) {
		fmt.Println("FreightTrain: arriving")
	} else {
		fmt.Println("FreightTrain: permission denied")
	}
}
func (p *FreightTrain) depart() {
	fmt.Println("FreightTrain: leaving station")
	p.mediator.notifyAboutDeparture()
}
func (p *FreightTrain) permitArrival() {
	fmt.Println("FreightTrain: arrival permited. Arriving")
	p.arrive()
}
