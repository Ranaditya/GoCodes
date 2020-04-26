package main

import (
	"fmt"
	"math"
)

type square struct {
	s float64
}

type circle struct {
	r float64
}

func (sq square) calcArea() float64 {
	return sq.s * sq.s
}
func (c circle) calcArea() float64 {
	return math.Pi * c.r * c.r
}

type shape interface {
	calcArea() float64
}

func info(sp shape) {
	fmt.Println(sp.calcArea())
}

func main() {
	s1 := square{
		s: 2.0,
	}
	c1 := circle{
		r: 3.0,
	}

	info(s1)
	info(c1)

}
