package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
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

	var n = map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map", n)

	var myGreeting = make(map[string]string)
	myGreeting["A"] = "Good morning."
	myGreeting["Jenny"] = "Bonjour."
	fmt.Println(myGreeting)

	//	HASH
	//res, err := http.Get("http://www-01.sil.org/linguistics/wordlists/english/wordlist/wordsEn.txt")
	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()

	buckets := make([][]string, 12)
	for i := 0; i < 12; i++ {
		buckets = append(buckets, []string{})
	}

	sc := bufio.NewScanner(res.Body)
	sc.Split(bufio.ScanWords)

	for sc.Scan() {
		word := sc.Text()
		key := HashBucket(word, 12)
		buckets[key] = append(buckets[key], word)
		//words[sc.Text()] = ""
	}

	for i := 0; i < 12; i++ {
		fmt.Println(i, " - ", len(buckets[i]))
	}
	fmt.Println(buckets[6])

	//if err := sc.Err(); err != nil {
	//	fmt.Fprintln(os.Stderr, "reading input:", err)
	//}
	//
	//i := 0
	//for k, _ := range words {
	//	fmt.Println(k)
	//	if i == 200 {
	//		break
	//	}
	//	i++
	//}

	//str := string(bs)
	//fmt.Println(str)

	letter := 'A'
	fmt.Printf("%v, %T \n", letter, letter)

	word := "Apple"
	fmt.Printf("%v, %T \n", word, word)
	letter = rune(word[0])
	fmt.Println(letter)

	GO := HashBucket("Go", 12)
	fmt.Println(GO)
}

// HashBucket hogehoge
func HashBucket(word string, buckets int) int {
	var sum int
	for _, v := range word {
		sum += int(v)
	}
	return sum % buckets
	//letter := int(word[0])
	//bucket := letter % buckets
	//return bucket
}
