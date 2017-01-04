package main

import "fmt"

func main() {
	f := factorialNormal(4)
	fmt.Println("Total:", f)

	//fmt.Println("Total: ", factorialChannel(4))
	c := factorialChannel(4)
	for n:= range c {
		fmt.Println("Total:", n)
	}

	in := gen()
	f2 := factorialPipeline(in)
	for n := range f2 {
		fmt.Println(n)
	}
}

func gen() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i< 10; i++ {
			for j := 3; j < 13; j++ {
				out <- j
			}
		}
		close(out)
	}()
	return out
}

func factorialPipeline(ints <-chan int) <-chan int {
	out := make(chan int)
	go func () {
		for n:= range ints{
			out <- factorialNormal(n)
		}
		close(out)
	}()
	return out
}

func factorialNormal(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}


// goroutine is helpful when you have many processing to do.
// if you want to run factorialNormal calculations for 1000 processes,
// then you want to utilize all of the CPU core using goroutine.
func factorialChannel(n int) chan int {
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