package main

func main() {
	var shape IShape = &Circle{x: 10, y: 5}
	var method IVisitor = &Perimetr{}
	shape.accept(method)
	shape = &Rectangle{x: 10, y: 5}
	shape.accept(method)
}
