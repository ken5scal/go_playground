package main

import (
	"fmt"
	"math"
	"sort"
)

type vehicles interface{}
type vehicle struct {
	Seats    int
	MaxSpeed int
	Color    string
}

type car struct {
	vehicle
	Wheels int
	Doors  int
}

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

// Value Receiver
func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

// pointer receiver
//func (c *Circle) area() float64 {
//	return c.radius * c.radius * math.Pi
//}

func info(z Shape) {
	// <- Interface Shape!
	fmt.Println(z)
	fmt.Println(z.area())
}

type IntSlice []int

func (p IntSlice) Len() int           { return len(p) }
func (p IntSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p IntSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main() {
	s := Square{10}
	info(s)

	c := Circle{10}
	info(c)
	info(&c)

	type peoplle []string
	studyGroup := peoplle{"Zeno", "john", "Al", "Jenny"}
	sort.Sort(sort.Reverse(sort.StringSlice(studyGroup)))
	fmt.Println(studyGroup)

	//s := []string{"Zeno", "john", "Al", "Jenny"}
	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
	slice := sort.IntSlice(n)
	slice.Sort()
	fmt.Println(slice)
	sort.Sort(sort.Reverse(slice))
	fmt.Println(slice)

	prius := car{}
	rides := []vehicles{prius}
	fmt.Println(rides)
}
