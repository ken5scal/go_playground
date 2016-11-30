package main

import (
	"math"
	"strings"
	"io"
	"golang.org/x/tour/reader"
	"time"
	"strconv"
	"sync"
	"fmt"
)

/*
	Interface
 */
// Golang Interface is implicit. You do not declare "implements"
// Can appear in any package without prearrangement.
type Abser interface {
	Abs() float64
}

type Vertex struct {
	X, Y float64
}

/*
	Ingeneral, all methos on given type should have either value/pointer receivers/
*/
// pointer type method
func (v *Vertex) Abs() float64 {
	if v == nil {
		// Handling nil receiver. this should be in method
		fmt.Println("<pointer t is nill>")
		return 0
	}
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

// value type method
func (v Vertex) AbsHoge() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

// Methods with pointer receivers can modify the value to which the receiver points
// pointer type method
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// Not recommended way to change type. It takes longer processing time
// Use value type method OR pointer type method
func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// Implements Stringer type interface
// fmt package look for this interface to PRINT values
func (v Vertex) String() string {
	return fmt.Sprintf("%v %v", v.X, v.Y)
}

func (v *Vertex) Error() string {
	return fmt.Sprint("Oh Error!:%v %v", v.X, v.Y)
}

type MyFloat float64
// You can only declare a mehtod with a receiver
// whose type is defined in hte same package as the method
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	v := Vertex{X:3, Y:4}

	var p *Vertex    // pointer to Vertex value
	p = &v //	& operator generates a pointer
	fmt.Println(*p)    // read through pointer. display {value of Struct}
	fmt.Println(p)    // should diplay &{value of Struct}
	fmt.Println((*p).X) // Formalized way of accessing Struct Field
	fmt.Println(p.X)    // Golang permits to do this way

	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v)) // AbsFunc(&v) will Fail
	v.Scale(10)    // Go interprets as (&v).Scale(10)
	// (&v).Scale(10) will be fine as well
	fmt.Println(v.Abs())
	fmt.Println((&v).Abs()) // Go interprets as (*(&v)).Abs()
	fmt.Println(v.AbsHoge())
	fmt.Println((&v).AbsHoge())
	ScaleFunc(&v, 10) //ScaleFunc(v, 10) fails
	fmt.Println(v.Abs())

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	/*
		Interfaces
	 */
	var a Abser
	f = MyFloat(-math.Sqrt2)
	a = f                          // <- Under the food, interface instance holds a value and concrete type
	// Concrete type is a type of struct assigned to iterface instance
	fmt.Printf("(%v, %T)\n", a, a) // <- Type is main.MyFloat
	fmt.Println(a.Abs())
	//a = v					// This will return error bc Vertex's Abs method is pointer type
	a = &v
	fmt.Printf("(%v, %T)\n", a, a) // <- Type is *main.Vertex
	fmt.Println(a.Abs())
	a = MyFloat(math.Ln10)
	fmt.Println(a.Abs())

	var a2 Abser
	fmt.Printf("(%v, %T)\n", a2, a2) // <- At this point, nil interface does not have any value and concrete type
	// a2.Abs()		will return Run-time error bc no implementation exists at this point

	var v2 *Vertex
	a2 = v2
	fmt.Printf("(%v, %T)\n", a2, a2)
	fmt.Println(a2.Abs())

	a2 = &Vertex{X:2, Y:3}
	fmt.Printf("(%v, %T)\n", a2, a2)
	fmt.Println(a2.Abs())

	// Empty interface that hold values of any type
	// Ex) fmt.Print takes any number of arguments of type interface{}
	var i interface{}
	fmt.Printf("(%v, %T)\n", i, i)  // <nil>, <nil>
	i = 42
	fmt.Printf("(%v, %T)\n", i, i)
	i = "hoge"
	fmt.Printf("(%v, %T)\n", i, i)

	// Type Assertion
	h := a2.(*Vertex)    // a2.(vertex) will fail
	fmt.Println(h)

	s, ok := i.(string)
	fmt.Println(s, ok)

	//f2 := i.(float64) // Panic will occur because interface is string type
	f2, ok := i.(float64) // by adding second output, you can avoid panic
	fmt.Println(f2)

	typeSwitch(&Vertex{X:2, Y:3})
	typeSwitch(MyFloat(math.Sqrt2))

	x := Vertex{3, 3}
	y := Vertex{1, 1}
	fmt.Println(x, y) // Vertex implements Stringer

	// Another example of Stringers
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}

	if err := runErrorVertex(); err != nil {
		fmt.Println(err)
	}

	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))

	r := strings.NewReader("Hello, Reader!")
	b := make([]byte, 8)
	for {
		n, err := r.Read(b) // Popoulates given byte slice of data
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			// End of stream
			break
		}
	}

	reader.Validate(MyReader{})

	// goroutine is a ligthweight threat
	// `go func()` starts a new goroutine with func()
	// Evaluation of func() is in the current go routine
	// but Execution will be in place of new goroutine
	// It runs in SAME memory space
	go say("world")
	say("hello")

	/*
		Concurrencies
	 */
	say := []int{7, 2, 8, -9, 4, 0}
	// Channel needs declaration
	c := make(chan int)
	// sends and receives block until the other side is ready
	// Synchronize gorouine
	go sum(say[:len(say) / 2], c) // sum number sis sent to channel
	go sum(say[len(say) / 2:], c)
	x1, y1 := <-c, <-c    // receives from channel
	fmt.Println(x1, y1, x1 + y1) // wait until both goroutines complete

	c = make(chan int, 2) // channel can be set buffer size
	c <- 1; c <- 2; // c <-3 Another c<-int will result in DEAD LOACK
	fmt.Println(<-c)
	fmt.Println(<-c)

	c = make(chan int, 5)
	fmt.Print("Fibonacci Ex: ")
	go fibonacci(cap(c), c)
	for i := range c {
		// i, ok := <- c
		// ok will be false if there are no more values to receive
		fmt.Print(strconv.Itoa(i) + " ")
	}

	fmt.Print("\nFibonacci2 Ex: ")
	c = make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Print(<-c)
		}
		quit <- 0
	}()
	fibonacci2(c, quit)

	// You don;t need to use channel if we just want to make sure
	// only one goruotine can access avariable at a time
	// -> MUTUAL EXCLUSION
	c2 := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c2.Inc("hoge")
	}
	time.Sleep(time.Millisecond)
	fmt.Println(c2.Value("hoge"))
}

type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()    // Lock so only one goroutine at a time can access here
	defer c.mux.Unlock()
	c.v[key]++
}

func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()    // Lock so only one goroutine at a time can access here
	defer c.mux.Unlock()
	return c.v[key]
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x + y
	}
	// Sender can close CHannel and notify that no more values
	// will be sent.
	close(c)
}

func fibonacci2(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		// Blocks until one of its cases can run
		// Chooses a channel at random if multiple are ready
		case c <- x:
			x, y = y, y + x
		case <-quit:
			fmt.Println("Quit")
			return
		default:
		// run if no other case is ready
		// can use to send/receive without blocking
		// in this case, . will be printed wheneve there is no other channel is ready
			fmt.Print(".")
		}
	}
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to channel
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

type IPAddr [4]byte
// Another example of Stringers
func (i IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", i[0], i[1], i[2], i[3])
}

func typeSwitch(a Abser) {
	switch v := a.(type) {    // switch case by the given concrete tpe
	case *Vertex:
		fmt.Printf("This is Vertex: %v \n", v)
	case MyFloat:
		fmt.Printf("This is float: %v \n", v)
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func runErrorVertex() error {
	return &Vertex{3, 3} //Implementing Error() by pointery type
}

// Example of Error type
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	// You need to convert e to float64(e) because
	// fmt.Sprintf(%v, e) will call e.Error and, inside Error(), ErrNegativeSqrt calls Sprintf(e),
	// So it results in infinite loop
	return fmt.Sprintf("cannot Sqrt negative num: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(-2)
	}
	return 0, nil
}

type MyReader struct{}

func (r MyReader) Read(b []byte) (int, error) {
	for x := range b {
		b[x] = 'A'
	}
	return len(b), nil
}

type PingPongPayload struct {
	Counter int
}
func ExamplePingPong() {

}