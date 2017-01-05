package main

import "fmt"

func main() {
	c := make(chan int)
	//c <- 1 Code Blocked here.
	// Channels are the pipes that connect concurrent goroutines. You can send values into channels from one goroutine and receive those values into another goroutine.
	go func() { c <- 1 }()
	fmt.Println(<-c)

	c1 := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c1 <- i
		}
		// Without closing, it will deadlock
		close(c1)
	}()

	// fmt.Println(i) This only print 0
	//for { // This will be deadlock
	for i := range c1 {
		fmt.Println(i)
	}
}
