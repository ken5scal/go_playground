package main

import "fmt"

func main() {
	f := factorial(4)
	fmt.Println("Total:", f)

	fmt.Println("Total: ", factorial2(4))
}

func factorial(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}

func factorial2(n int) int {
	c := make(chan int)
	total := 1
	go func() {
		for i := n; i > 0; i-- {
			c<-i
		}
		close(c)
	}()

	for i := range c {
		total *= i
	}
	return total
}