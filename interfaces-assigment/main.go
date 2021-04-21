package main

import (
	"fmt"
	"math"
)

type shape interface {
	Name() string
	getArea() float64
}
type Triangle struct {
	Base   float64
	Height float64
}
type Square struct {
	Side float64
}

func (t Triangle) Name() string {
	return "Triangle"
}

func (t Triangle) getArea() float64 {
	return (t.Base * t.Height) / 2
}

func (s Square) Name() string {
	return "Square"
}

func (s Square) getArea() float64 {
	return math.Sqrt(s.Side)
}

func main() {
	t := Triangle{Base: 10, Height: 3}
	s := Square{Side: 4}

	printArea(t)
	printArea(s)
}

func printArea(s shape) {
	fmt.Printf("Area of %s is %.2f\n", s.Name(), s.getArea())
}
