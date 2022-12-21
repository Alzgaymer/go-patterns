package main

type Ak47 struct {
	Gun
}
type Musket struct {
	Gun
}

func newAk47() IGun {
	return Ak47{
		Gun: Gun{
			Name:   "AK47",
			Damage: 20,
		},
	}
}
func newMusket() IGun {
	return Musket{
		Gun: Gun{
			Name:   "Musket",
			Damage: 10,
		},
	}
}
