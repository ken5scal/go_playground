package main

import "fmt"

func main() {
	f := factorial(4)
	fmt.Println("Total:", f)

	//fmt.Println("Total: ", factorial2(4))
	c := factorial2(4)
	for n := range c {
		fmt.Println("Total:", n)
	}
}

func factorial(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}

// goroutine is helpful when you have many processing to do.
// if you want to run factorial calculations for 1000 processes,
// then you want to utilize all of the CPU core using goroutine.
func factorial2(n int) chan int {
	c := make(chan int)

	//total := 1
	//go func() {
	//	for i := n; i > 0; i-- {
	//		c<-i
	//	}
	//	close(c)
	//}()
	//
	//for i := range c {
	//	total *= i
	//}
	//return total

	go func() {
		total := 1
		for i := n; i > 0; i-- {
			total *= i
		}
		c <- total
		close(c)
	}()
	return c
}
