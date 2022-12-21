package main

import "fmt"

type IObserver interface {
	notify(string)
	getEmail() string
}

type Member struct {
	email string
}

func (m *Member) notify(itemName string) {
	fmt.Printf("Sending email to customer %s for item %s\n", m.email, itemName)
}
func (m *Member) getEmail() string {
	return m.email
}
