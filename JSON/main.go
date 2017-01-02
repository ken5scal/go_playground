package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Person struct {
	First  string
	Last string	`json:"-"`
	Age int	`json:"wisdom score"`
	notExported int
}

func main() {
	fmt.Println("----Marshaling(string -> slice of byte)-----")
	p1 := Person{"A", "B", 20, 0}
	bs, _ := json.Marshal(&p1)
	fmt.Println(bs)
	fmt.Printf("%T \n", bs)
	fmt.Println(string(bs))


	fmt.Println("----Unmarshaling-----")
	var p2 Person
	bs = []byte(`{"First":"James", "Last":"bond", "wisdom score":20}`)
	json.Unmarshal(bs, &p2)
	fmt.Println(p2.First)
	fmt.Println(p2.Last)
	fmt.Println(p2.Age)
	fmt.Printf("%T \n", p2)

	fmt.Println("----Encode(stream)-----")
	p3 := Person{"A", "B", 20, 0}
	json.NewEncoder(os.Stdout).Encode(p3)

	fmt.Println("----Decode(stream)-----")
	var p4 Person
	rdr := strings.NewReader(`{"First":"James", "Last":"bond", "wisdom score":20}`)
	json.NewDecoder(rdr).Decode(&p4)
	fmt.Println(p4.Age)
}
