package main

import (
	"math"
	"fmt"
)


/*
	Methods
 */
type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

type MyFloat float64
// You can only declare a mehtod with a receiver
// whose type is defined in hte same package as the method
func (f MyFloat) Abs() float64 {
	if f<0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	v := Vertex{X:3, Y:4}
	fmt.Println(v.Abs())
	fmt.Println(Abs(v))

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
