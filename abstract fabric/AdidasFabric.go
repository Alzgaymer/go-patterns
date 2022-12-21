package main

type Adidas struct {
}

func (a *Adidas) makeShoe(s string, ss int) IShoe {
	return &AdidasShoe{Shoe: Shoe{Logo: s, Size: ss}}
}
func (a *Adidas) makeShirt(s string, ss int) IShirt {
	return &AdidasShirt{Shirt: Shirt{Logo: s, Size: ss}}
}
