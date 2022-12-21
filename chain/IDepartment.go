package main

type IDepartment interface {
	execute(*Patient)
	setNext(IDepartment)
}
