package main

type Object struct {
	state string
}

func (o *Object) set(s string) {
	o.state = s
}
func (o *Object) get() string {
	return o.state
}

func (o *Object) create() *Snap {
	return &Snap{
		state: o.state,
	}
}

func (o *Object) restore(s *Snap) {
	o.state = s.state
}
