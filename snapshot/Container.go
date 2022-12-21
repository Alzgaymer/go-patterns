package main

type Container struct {
	arr []*Snap
}

func (c *Container) append(s *Snap) {
	c.arr = append(c.arr, s)
}
func (c *Container) getSnap(index int) *Snap {
	return c.arr[index]
}
