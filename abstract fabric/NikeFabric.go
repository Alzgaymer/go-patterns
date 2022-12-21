package main

type Nike struct {
}

func (a *Nike) makeShoe(s string, ss int) IShoe {
	return &NikeShoe{Shoe: Shoe{Logo: s, Size: ss}}
}
func (a *Nike) makeShirt(s string, ss int) IShirt {
	return &NikeShirt{Shirt: Shirt{Logo: s, Size: ss}}
}
