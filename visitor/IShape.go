package main

type IShape interface {
	accept(IVisitor)
	getType() string
}
