package main

import "fmt"

var x = 42 // package level scope

func main() {
	fmt.Println(x)
	foo()
	//y:=17 <- block level
}

func foo() {
	fmt.Println(x)
}
