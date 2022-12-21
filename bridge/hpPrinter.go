package main

import "fmt"

type HP struct {
}

func (h *HP) PrintMsg() {
	fmt.Println("Printing from HP printer")
}
