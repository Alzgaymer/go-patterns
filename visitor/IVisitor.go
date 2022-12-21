package main

type IVisitor interface {
	VisitorForCircle(*Circle)
	VisitorForRectengle(*Rectangle)
}
