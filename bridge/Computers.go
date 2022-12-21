package main

import "fmt"

type Mac struct {
	printer IPrinter
}

func (m *Mac) SetPrinter(p IPrinter) {
	m.printer = p
}
func (m *Mac) Print() {
	fmt.Println("Printing from Mac")
	m.printer.PrintMsg()
}
