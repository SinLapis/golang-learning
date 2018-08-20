package main

import "fmt"

func compare(a int, b int) (small int, big int) {
	if a > b {
		small, big = b, a
	} else {
		small, big = a, b
	}
	return
}

func main() {
	a, b := 7, 2
	a, b = compare(a, b)
	c, d := 3, 4
	c, d = compare(c, d)
	fmt.Printf("%d, %d\n", a, b)
	fmt.Printf("%d, %d\n", c, d)
}
