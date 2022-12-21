package main

type Shoe struct {
	Logo string
	Size int
}
type Shirt struct {
	Logo string
	Size int
}

type IShirt interface {
	GetLogo() string
	GetSize() int
}
type IShoe interface {
	GetLogo() string
	GetSize() int
}

func (s *Shoe) GetLogo() string {
	return s.Logo
}
func (s *Shoe) GetSize() int {
	return s.Size
}
func (s *Shirt) GetLogo() string {
	return s.Logo
}
func (s *Shirt) GetSize() int {
	return s.Size
}
