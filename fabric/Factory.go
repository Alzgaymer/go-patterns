package main

import "fmt"

func getGun(str string) (IGun, error) {

	if str == "AK47" {
		return newAk47(), nil
	} else if str == "Musket" {
		return newMusket(), nil
	}
	return nil, fmt.Errorf("can`t create a gun")
}
