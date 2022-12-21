package main

import "fmt"

func main() {
	con := Container{
		arr: make([]*Snap, 0),
	}
	o := Object{state: "1"}
	con.append(o.create())
	fmt.Printf("Object current state: %s\n", o.get())

	o.set("2")
	con.append(o.create())
	fmt.Printf("Object current state: %s\n", o.get())

	o.set("3")
	con.append(o.create())
	fmt.Printf("Object current state: %s\n", o.get())
	o.restore(con.getSnap(1))
	fmt.Printf("Object current state: %s\n", o.get())
	o.restore(con.getSnap(0))
	fmt.Printf("Object current state: %s\n", o.get())
}
