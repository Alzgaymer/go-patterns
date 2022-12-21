package main

type Reactengle struct {
	X, Y  int
	Color string
}

func (f *Reactengle) clone() IShape {
	return &Reactengle{
		X:     f.X,
		Y:     f.Y,
		Color: f.Color,
	}
}
