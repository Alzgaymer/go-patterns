package main

type Circle struct {
	X, Y  int
	Color string
	D     float64
}

func (c *Circle) clone() IShape {
	return &Circle{
		X:     c.X,
		Y:     c.Y,
		D:     c.D,
		Color: c.Color,
	}
}
