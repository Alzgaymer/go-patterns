package main

func main() {
	comp := Mac{}
	hp := &HP{}
	comp.SetPrinter(hp)

	comp.Print()
}
