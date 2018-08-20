package main

import "fmt"

func floatAva(nums []float64) {
	sum := float64(0)
	for _, v := range nums {
		sum += v
	}
	fmt.Printf("%f\n", sum/float64((len(nums))))
	return
}

func main() {
	floatAva([]float64{1.2, 3.5, 6.43, 9.03965})
}
