package main

import "fmt"

func main() {
	c := make(chan int)
	//c <- 1 Code Blocked here.
	// Channels are the pipes that connect concurrent goroutines. You can send values into channels from one goroutine and receive those values into another goroutine.
	go func() {c<-1}()
	fmt.Println(<-c)
}
