package main

type ITrain interface {
	arrive()
	depart()
	permitArrival()
}
