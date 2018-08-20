package main

import "fmt"

func plus2(a int) (b int) {
	return a + 2
}

func lenPlus1(str string) (al int) {
	return len(str) + 1
}

func simpleMap(f func(int) int, args []int) (results []int) {
	results = make([]int, len(args))
	for k, v := range args {
		results[k] = f(v)
	}
	return
}

func main() {
	l1 := []int{1, 3, 5, 11, 45}
	fmt.Printf("%v\n", simpleMap(plus2, l1))
}
