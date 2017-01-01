package main

import (
	"fmt"
	"net/http"
	"log"
	"bufio"
	"os"
)

func main() {
	// reference type. pointing to actual data type
	//
	m := make(map[string]int)
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

	//	HASH
	res, err := http.Get("http://www-01.sil.org/linguistics/wordlists/english/wordlist/wordsEn.txt")
	if err != nil {
		log.Fatalln(err)
	}

	words := make(map[string]string)
	sc := bufio.NewScanner(res.Body)
	sc.Split(bufio.ScanWords)
	for sc.Scan() {
		words[sc.Text()] = ""
	}

	if err := sc.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	i := 0
	for k, _ := range words {
		fmt.Println(k)
		if i == 200 {
			break
		}
		i++
	}

	//str := string(bs)
	//fmt.Println(str)
}
