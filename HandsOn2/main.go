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

func (p person) pSpeak() {
	fmt.Println(p.fname, p.lname, "says how r u doing?")
}

func (sa secretagent) saSpeak() {
	fmt.Println("secret agent", sa.p.fname, sa.p.lname, "has l2kill", sa.l2kill)
}

func main() {
	p := person{
		"John", "Snow",
	}
	s := secretagent{
		person{"James", "Bond"},
		true,
	}
	p.pSpeak()
	s.saSpeak()
	s.p.pSpeak()
}
