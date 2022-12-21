package main

import "fmt"

type ISubject interface {
	register(IObserver)
	deregister(int)
	notifyAll()
}
type Item struct {
	customers []IObserver
	name      string
	inStock   bool
}

func newItem(name string) *Item {
	return &Item{
		name: name,
	}
}
func (i *Item) updateAvailability() {
	fmt.Printf("Item %s is in stock right now!\n", i.name)
	i.inStock = true
	i.notifyAll()
}
func (i *Item) register(o ...IObserver) {
	i.customers = append(i.customers, o...)
}
func (i *Item) deregister(in int) {
	i.customers = append(i.customers[:in], i.customers[in+1:]...)
}
func (i *Item) notifyAll() {
	for _, v := range i.customers {
		v.notify(i.name)
	}
}
