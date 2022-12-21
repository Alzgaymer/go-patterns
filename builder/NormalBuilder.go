package main

/*
type IBuilder interface {
	setWalls()
	setDoors()
	GetResult() House
}
*/

type NormalBuilder struct {
	Walls int
	Doors int
}

func (s *NormalBuilder) setWalls() {
	s.Walls = 4
}

func (s *NormalBuilder) setDoors() {
	s.Doors = 1
}
func (s *NormalBuilder) GetResult() House {
	return House{
		Walls: s.Walls,
		Doors: s.Doors,
	}
}

func getNormalBuilder() *NormalBuilder {
	return &NormalBuilder{}
}
