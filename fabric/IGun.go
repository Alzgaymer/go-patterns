package main

type IGun interface {
	GetDamage() float32
	GetName() string
	String() string
}
