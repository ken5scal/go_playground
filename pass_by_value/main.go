package main

import "fmt"

func main() {
	age := 44
	changeMe(age)
	fmt.Println(age) //44 <- no CHange
	changeMeReal(&age)
	fmt.Println(age) //24

	m := make(map[string]int) //making references to data
	changeMeNow(m)
	fmt.Println(m["Todd"]) //44
}

func changeMe(z int) {
	z = 24
}

func changeMeReal(z *int) {
	*z = 24
}

func changeMeNow(z map[string]int) {
	z["Todd"] = 44
}
