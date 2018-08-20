package main

import (
	"fmt"
	"unicode/utf8"
)

func printA() {
	fmt.Printf("printA-----start-----\n")
	for i := 0; i < 100; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("A")
		}
		fmt.Printf("\n")
	}
	fmt.Printf("printA------end------\n")
}

func countChar() {
	fmt.Printf("countChar---start----\n")
	s := "asSASA ddd dsjkdsjs dk"
	s2 := "asSabc ddd dsjkdsjs dk"
	fmt.Printf("bytes: %d\n", len(s))
	fmt.Printf("chars: %d\n", utf8.RuneCountInString(s))
	fmt.Printf("bytes: %d\n", len(s2))
	fmt.Printf("chars: %d\n", utf8.RuneCountInString(s2))
	fmt.Printf("countChar----end-----\n")
}

func evertString() {
	fmt.Printf("evertString--start---\n")
	s := "foobar"
	l := len(s)
	r := make([]byte, l)
	for i := 0; i < l; i++ {
		r[i] = s[l-i-1]
	}
	fmt.Printf("%s\n", r)
	fmt.Printf("evertString---end----\n")
}

func main() {
	printA()
	countChar()
	evertString()
	floatsAva()
}
