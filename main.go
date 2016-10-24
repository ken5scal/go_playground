package main

import (
	"math"
	"fmt"
)

/*
	Interface
 */
// Golang Interface is implicit. You do not declare "implements"
// Can appear in any package without prearrangement.
type Abser interface {
	Abs() float64
}

type Vertex struct {
	X, Y float64
}

/*
	Ingeneral, all methos on given type should have either value/pointer receivers/
*/
// pointer type method
func (v *Vertex) Abs() float64 {
	if v == nil {
		// Handling nil receiver. this should be in method
		fmt.Println("<pointer t is nill>")
		return 0
	}
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

// value type method
func (v Vertex) AbsHoge() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

// Methods with pointer receivers can modify the value to which the receiver points
// pointer type method
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// Not recommended way to change type. It takes longer processing time
// Use value type method OR pointer type method
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

	var p *Vertex    // pointer to Vertex value
	p = &v //	& operator generates a pointer
	fmt.Println(*p)    // read through pointer. display {value of Struct}
	fmt.Println(p)    // should diplay &{value of Struct}
	fmt.Println((*p).X) // Formalized way of accessing Struct Field
	fmt.Println(p.X)    // Golang permits to do this way

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
	a = f                          // <- Under the food, interface instance holds a value and concrete type
	// Concrete type is a type of struct assigned to iterface instance
	fmt.Printf("(%v, %T)\n", a, a) // <- Type is main.MyFloat
	fmt.Println(a.Abs())
	//a = v					// This will return error bc Vertex's Abs method is pointer type
	a = &v
	fmt.Printf("(%v, %T)\n", a, a) // <- Type is *main.Vertex
	fmt.Println(a.Abs())
	a = MyFloat(math.Ln10)
	fmt.Println(a.Abs())

	var a2 Abser
	fmt.Printf("(%v, %T)\n", a2, a2) // <- At this point, nil interface does not have any value and concrete type
	// a2.Abs()		will return Run-time error bc no implementation exists at this point

	var v2 *Vertex
	a2 = v2
	fmt.Printf("(%v, %T)\n", a2, a2)
	fmt.Println(a2.Abs())

	a2 = &Vertex{X:2, Y:3}
	fmt.Printf("(%v, %T)\n", a2, a2)
	fmt.Println(a2.Abs())

	// Empty interface that hold values of any type
	// Ex) fmt.Print takes any number of arguments of type interface{}
	var i interface{}
	fmt.Printf("(%v, %T)\n", i, i)  // <nil>, <nil>
	i = 42
	fmt.Printf("(%v, %T)\n", i, i)
	i = "hoge"
	fmt.Printf("(%v, %T)\n", i, i)

	// Type Assertion
	h := a2.(*Vertex)    // a2.(vertex) will fail
	fmt.Println(h)

	s, ok := i.(string)
	fmt.Println(s, ok)

	//f2 := i.(float64) // Panic will occur because interface is string type
	f2, ok := i.(float64) // by adding second output, you can avoid panic
	fmt.Println(f2)

	typeSwitch(&Vertex{X:2, Y:3})
	typeSwitch(MyFloat(math.Sqrt2))
}


func typeSwitch(a Abser) {
	switch v := a.(type) {	// switch case by the given concrete tpe
	case *Vertex:
		fmt.Println("This is Vertex: %v", v)
	case MyFloat:
		fmt.Print("This is float: %f", v)
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}