package main

import "fmt"

func fibonacci(n int) {
	a, b := 1, 1
	fmt.Printf("%d %d ", a, b)
	var c int
	for i := 2; i < n; i++ {
		c = a + b
		fmt.Printf("%d ", c)
		a = b
		b = c
	}
	fmt.Printf("\n")
}

func main() {
	fibonacci(5)
	fibonacci(10)
}
