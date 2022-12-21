package main

type Circle struct {
	x, y float32
}

func (c *Circle) accept(v IVisitor) {
	v.VisitorForCircle(c)
}
func (c *Circle) getType() string {
	return "Circle"
}
