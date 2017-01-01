package main

import "fmt"

func main() {
	// reference type. pointing to actual data type
	//
	m := make(map[string ]int)
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)
	delete(m, "k2")
	fmt.Println("map:", m)

	v, ok := m["k1"]
	fmt.Println("ok?:", ok, v)

	var n = map[string]int{"foo" : 1, "bar":2}
	fmt.Println("map", n)

	var myGreeting = make(map[string]string)
	myGreeting["A"] = "Good morning."
	myGreeting["Jenny"] = "Bonjour."
	fmt.Println(myGreeting)
}
