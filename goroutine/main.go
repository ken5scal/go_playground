package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup
var counter int
var counterATOM int64
var mutex sync.Mutex

func main() {
	// Concurrency
	go fuga()
	// This it self wont finish bc main goroutine finishes right off
	// SO Add WaitGroup!

	wg.Add(4)
	go foo()
	go bar()

	go incrementor("Foo:")
	go incrementor("Bar:")
	wg.Wait()
}

func fuga() {
	for i := 0; i < 10000; i++ {
		fmt.Println("Fuga:", i)
	}
}
func bar() {
	for i := 0; i < 50; i++ {
		fmt.Println("Bar:", i)
	}
	wg.Done()
}
func foo() {
	for i := 0; i < 50; i++ {
		fmt.Println("Foo:", i)
	}
	wg.Done()
}

func incrementor(s string) {
	for i := 0; i < 20; i++ {
		// % go run -race main.go
		// <- Found 1 data race(s) (exit status 66) : Race Condition!
		// So DO MUTEX(Only one person can access to the goroutine <- atomicity)
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)

		// Mutex Pattern1
		atomic.AddInt64(&counterATOM, 1)

		// Mutex Pattern2(Repetittive way)
		mutex.Lock()
		x := counter
		x++
		counter = x
		// Or just counter ++
		fmt.Println(s, i, "Counter: ", counter, "CounterATOM: ", counterATOM)
		mutex.Unlock()
	}
	wg.Done()
}
