package main

type Director struct {
	builder IBuilder
}

func (d *Director) Build() House {
	d.builder.setDoors()
	d.builder.setWalls()
	return d.builder.GetResult()
}
func newDirector(b IBuilder) *Director {
	return &Director{
		builder: b,
	}
}
