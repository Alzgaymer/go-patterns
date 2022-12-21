package main

type Rectangle struct {
	x, y float32
}

func (r *Rectangle) accept(v IVisitor) {
	v.VisitorForRectengle(r)
}
func (r *Rectangle) getType() string {
	return "Rectangle"
}
