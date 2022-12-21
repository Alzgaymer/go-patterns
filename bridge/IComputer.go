package main

type IComputer interface {
	SetPrinter(IPrinter)
	Print()
}
