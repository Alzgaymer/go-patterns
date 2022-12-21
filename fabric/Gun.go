package main

import "fmt"

type Gun struct {
	Name   string
	Damage float32
}

func (g Gun) GetName() string {
	return g.Name
}
func (g Gun) GetDamage() float32 {
	return g.Damage
}
func (g Gun) String() string {
	return fmt.Sprintf("Gun name:%s, it`s damage:%v", g.Name, g.Damage)
}
