package main

type Snap struct {
	state string
}

func (s *Snap) getState() string {
	return s.state
}
