package main

import (
	"fmt"
	"sync"
)

func main() {
	in := gen(2, 3)

	// Fan Out
	c1 := sq(in)
	c2 := sq(in)
	
	for n := range merge(c1, c2) {
		fmt.Println(n) // 4 then 9, or 9 then 4
	}
}
func merge(cs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(cs))

	for _, c := range cs {
		go func(ch chan int) {
			for n := range ch {
				out <- n
			}
		}(c)
	}
	return out
}

func gen(nums ...int) <-chan int {
	fmt.Printf("Type of nums %T\n", nums)
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(nums <-chan int) <- chan int {
	out := make(chan int)
	go func() {
		for _, num := range nums {
			out <- num * num
		}
	}()
	return out
}
