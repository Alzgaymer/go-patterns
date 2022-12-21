package main

import (
	"fmt"
	"math"
)

type Perimetr struct {
}

func (p *Perimetr) VisitorForCircle(c *Circle) {
	fmt.Printf("Perimetr of %s: %v\n", c.getType(), 2*math.Pi*c.x)
}
func (p *Perimetr) VisitorForRectengle(r *Rectangle) {
	fmt.Printf("Perimetr of %s: %v\n", r.getType(), 2*r.x+2*r.y)
}
