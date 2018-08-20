package main

import "fmt"

func normal() {
	fmt.Printf("normal-----start-----\n")
	for i := 0; i < 10; i++ {
		fmt.Printf("%d\n", i)
	}
	fmt.Printf("normal------end------\n")
}

func noFor() {
	fmt.Printf("noFor------start-----\n")
	i := 0
repeat:
	fmt.Printf("%d\n", i)
	if i < 9 {
		i++
		goto repeat
	}
	fmt.Printf("noFor-------end------\n")
}

func traverseArray() {
	fmt.Printf("traverseArray-start--\n")
	l := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, i := range l {
		fmt.Printf("%d\n", i)
	}
	fmt.Printf("traverseArray--end---\n")
}

func main() {
	normal()
	noFor()
	traverseArray()
	return
}
