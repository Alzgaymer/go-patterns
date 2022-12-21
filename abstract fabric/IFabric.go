package main

type IFabric interface {
	makeShoe(string, int) IShoe
	makeShirt(string, int) IShirt
}

func GetFabric(s string) IFabric {
	switch s {
	case "Adidas":
		return &Adidas{}
	case "Nike":
		return &Nike{}
	}
	return nil
}
