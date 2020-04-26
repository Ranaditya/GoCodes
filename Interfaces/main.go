package main

import (
	"fmt"
)

type engbot struct{}
type spbot struct{}

type bot interface {
	getGreetings() string
}

func main() {

	e := engbot{}
	s := spbot{}

	printGreetings(e)
	printGreetings(s)
}

func (engbot) getGreetings() string {
	return "Hi there!"
}

func (spbot) getGreetings() string {
	return "Hola!"
}

func printGreetings(b bot) {
	fmt.Println(b.getGreetings())
}
