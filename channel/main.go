package main

import (
	"fmt"
	"time"
)

func main() {
	// Do not communicate by sharing memory (Dont block by mutex or atomic function when goroutine accesses to it),
	// share memory by communication

	// ANd this is what channel does
	c := make(chan int) // Unbuffered channel

	go func() {
		for i := 0; i < 10; i++ {
			c <- i // blocks until fmt.println(<-c) is executed
		}
	}()

	go func() {
		for {
			fmt.Println(<-c)
		}
	}()

	time.Sleep(time.Second) // Gives time to Execute all code

	// ----- Better because it wont need additional goroutine, and sleep function"
	//c2 := make(chan int)
	//go func() {
	//	for i := 0; i < 10; i++ {
	//		c2 <- i
	//		fmt.Printf("i: %v\n", i)
	//	}
	//	close(c2) // Channel is closed, you cannot put values to channel.
	//}()
	//
	//for n := range c2 { // Wait until it receives.
	//	fmt.Println(n)
	//}
}
