package main

import "fmt"

func plusTwo() (f func(int) int) {
	f = func(a int) (b int) {
		b = a + 2
		return
	}
	return
}

func plusX(x int) (f func(int) int) {
	f = func(a int) (b int) {
		b = a + x
		return
	}
	return
}

func main() {
	p := plusTwo()
	p5 := plusX(5)
	fmt.Printf("puls 2: %d\n", p(2))
	fmt.Printf("plus 5: %d\n", p5(3))
}
