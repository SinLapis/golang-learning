package main

import "fmt"

func floatsAva() {
	fmt.Printf("floatsAva---start----\n")
	floats := []float64{1.1, 2.3, 4.6, 8.77, 9.564}
	sum := float64(0.0)
	for _, f := range floats {
		sum += f
	}
	fmt.Printf("%f\n", sum/float64(len(floats)))
	fmt.Printf("floatsAva----end-----\n")
}

func main() {
	floatsAva()
}
