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

/*
	Ingeneral, all methos on given type should have either value/pointer receivers/
*/
// Value method
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

// Pointer receivers method
func (v *Vertex) AbsHoge() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func AbsFunc(v Vertex) float64 {
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

type Abser interface {
	Abs() float64
}

func main() {
	v := Vertex{X:3, Y:4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v)) // AbsFunc(&v) will Fail
	v.Scale(10)    // Go interprets as (&v).Scale(10)
	// (&v).Scale(10) will be fine as well
	fmt.Println(v.Abs())
	fmt.Println((&v).Abs()) // Go interprets as (*(&v)).Abs()
	fmt.Println(v.AbsHoge())
	fmt.Println((&v).AbsHoge())
	ScaleFunc(&v, 10) //ScaleFunc(v, 10) fails
	fmt.Println(v.Abs())

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	/*
		Interfaces
	 */
	var a Abser
	f = MyFloat(-math.Sqrt2)
	a = f
	fmt.Println(a.Abs())
}
