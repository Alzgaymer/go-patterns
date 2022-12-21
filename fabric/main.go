package main

import "fmt"

func main() {
	a, _ := getGun("AK47")
	fmt.Println(a)
	_, err := getGun("Jopa")
	if err != nil {
		fmt.Print(err)
	}
}
