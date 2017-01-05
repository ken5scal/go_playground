package main

import (
	"fmt"
	"sync"
)

func main() {
	Nto1()
}

func Nto1() {
	// Nto1
	// Classic Way
	c3 := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		//wg.Add(1) // This is bad, because goroutine is trying to access shared variable
		for i := 0; i < 10; i++ {
			c3 <- i
		}
		wg.Done()
	}()

	go func() {
		//wg.Add(1) // This is bad, because goroutine is trying to access shared variable
		for i := 0; i < 10; i++ {
			c3 <- i
		}
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(c3)
	}()

	for n := range c3 {
		fmt.Println(n)
	}

	// semapho pattern: Something like flag, tells program what to do.

	n := 10
	c := make(chan int)     // Unbuffered channel
	done := make(chan bool) // Semapho: indicates loop is done

	for i := 0; i < n; i++ {
		go func() {
			for i := 0; i < 10; i++ {
				c <- i // blocks until fmt.println(<-c) is executed
			}
			done <- true
		}()
	}

	//
	//go func() {
	//	for i := 0; i < 10; i++ {
	//		c <- i
	//	}
	//	done <- true
	//}()

	go func() {
		//<-done
		//<-done
		for i := 0; i < n; i++ {
			<-done
		}
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}
