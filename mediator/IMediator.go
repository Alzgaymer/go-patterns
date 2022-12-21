package main

type IMediator interface {
	canArrive(ITrain) bool
	notifyAboutDeparture()
}
