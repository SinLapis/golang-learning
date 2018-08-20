package main

import "fmt"

func ambiguousNums(nums ...int) {
	fmt.Printf("start\n")
	for _, v := range nums {
		fmt.Printf("%d\n", v)
	}
	fmt.Printf("end\n")
}

func main() {
	ambiguousNums(1, 2, 3, 4, 5, 6)
	ambiguousNums(12, 35, 66)
}
