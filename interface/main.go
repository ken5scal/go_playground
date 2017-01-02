package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Square struct {
	side float64
}

// Concrete type implementing interface
func (z Square) area() float64 {
	return z.side * z.side
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func info(z Shape) {
	// <- Interface Shape!
	fmt.Println(z)
	fmt.Println(z.area())
}

func main() {
	s := Square{10}
	info(s)

	c := Circle{10}
	info(c)
}
