package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	// Concurrency
	go fuga() // This it self wont finish bc main goroutine finishes right off

	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}
func fuga() {
	for i := 0; i < 100000000000; i++ {
		fmt.Println("Fuga:",i )
	}
}
func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("Bar:",i )
	}
	wg.Done()
}
func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("Foo:",i )
	}
	wg.Done()
}
