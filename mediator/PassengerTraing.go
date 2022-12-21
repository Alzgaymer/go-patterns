package main

import "fmt"

type PassengerTrain struct {
	mediator IMediator
}

func (p *PassengerTrain) arrive() {
	if p.mediator.canArrive(p) {
		fmt.Println("PassengerTrain: arriving")
	} else {
		fmt.Println("PassengerTrain: permission denied")
	}
}
func (p *PassengerTrain) depart() {
	fmt.Println("PassengerTrain: leaving station")
	p.mediator.notifyAboutDeparture()
}
func (p *PassengerTrain) permitArrival() {
	fmt.Println("PassengerTrain: arrival permited. Arriving")
	p.arrive()
}
