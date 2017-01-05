package main

import (
	"fmt"
	"sync"
)

/*
Challenge #1:
-- Change the code above to execute 1,000 factorial computations concurrently and in parallel
-- Use the FAN-OUT/FAN_IN pattern to accomplish this

Challenge #2:
-- While running the factorial computations, try to find how much of your resources are being used.
*/
func main() {
	in := gen()
	c1 := factorial(in) //<- worker
	c2 := factorial(in) //<- worker
	c3 := factorial(in) //<- worker
	for n := range merge(c1, c2, c3) {
		fmt.Println(n)
	}
}

func gen() chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			for j := 3; j < 13; j++ {
				out <- j
			}
		}
		close(out)
	}()
	return out
}

func factorial(nums chan int) chan int {
	out := make(chan int)

	go func() {
		for num := range nums {
			out <- func(n int) int {
				total := 1
				for i := n; i > 0; i-- {
					total *= i
				}
				return total
			}(num)
		}
		close(out)
	}()

	return out
}

func merge(cs ...chan int) chan int {
	out := make(chan int)

	var wg sync.WaitGroup
	wg.Add(len(cs))

	for _, c := range cs {
		go func(ch chan int) {
			for n := range ch {
				out <- n
			}
			wg.Done()
		}(c)
	}

	// Close out once all the output goroutine are done.
	// This must start after the wg.Add Call
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
