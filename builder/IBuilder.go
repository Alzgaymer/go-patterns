package main

type IBuilder interface {
	setWalls()
	setDoors()
	GetResult() House
}

func getBuilder(s string) IBuilder {
	switch s {
	case "normal":
		return getNormalBuilder()
	}
	return nil
}
func (d *Director) setBuilder(b IBuilder) {
	d.builder = b
}
