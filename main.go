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

// Methods with pointer receivers can modify the value to which the receiver points
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

type MyFloat float64
// You can only declare a mehtod with a receiver
// whose type is defined in hte same package as the method
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	v := Vertex{X:3, Y:4}
	fmt.Println(v.Abs())
	fmt.Println(Abs(v))
	v.Scale(10) 	// Go interprets as (&v).Scale(10)
	fmt.Println(v.Abs())
	ScaleFunc(&v, 10) //ScaleFunc(v, 10) fails
	fmt.Println(v.Abs())

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
