package main

import (
	"fmt"
)

type person struct {
	fname string
	lname string
}
type secretagent struct {
	p      person
	l2kill bool
}

type human interface {
	speak()
}

func (p person) speak() {
	fmt.Println(p.fname, p.lname, "says how r u doing?")
}

func (sa secretagent) speak() {
	fmt.Println("secret agent", sa.p.fname, sa.p.lname, "has l2kill", sa.l2kill)
}

func printHuman(h human) {
	h.speak()
}

func main() {
	p := person{
		"John", "Snow",
	}
	s := secretagent{
		person{"James", "Bond"},
		true,
	}
	printHuman(p)
	printHuman(s)
	printHuman(s.p)
}
