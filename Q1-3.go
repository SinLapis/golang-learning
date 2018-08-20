package main

import "fmt"

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

func

func main() {
	printA()
}
