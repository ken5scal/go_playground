package main

import "fmt"

func main() {
	//2
	//var x int <- no longer package level scope
	//increment := func() int {  <- anonymous function
	//	x++
	//	return x
	//}
	increment := wrapper()
	fmt.Println(increment())
	fmt.Println(increment())
}

func wrapper() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

// 1
// var x int <- package level scope
//func increment() int {
//	x++
//	return x
//}
