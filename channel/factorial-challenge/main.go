package main

import "fmt"

func main() {
	f := factorial(4)
	fmt.Println("Total:", f)
}

func factorial(n int) interface{} {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}
