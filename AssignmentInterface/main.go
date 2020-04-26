package main

import (
	"fmt"
)

type shape interface {
	getArea() float64
}

type triangle struct {
	base float64
	ht   float64
}

type square struct {
	side float64
}

func main() {
	t := triangle{
		base: 2,
		ht:   4,
	}
	s := square{
		side: 2,
	}
	printArea(t)
	printArea(s)
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.ht
}

func (s square) getArea() float64 {
	return s.side * s.side
}

func printArea(s shape) {
	fmt.Println("Area of shape", s.getArea())
}
