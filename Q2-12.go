package main

import "fmt"

func simpleMin(nums []int) (minNum int) {
	minNum = nums[0]
	for _, v := range nums[1:] {
		if v < minNum {
			minNum = v
		}
	}
	return
}

func simpleMax(nums []int) (maxNum int) {
	maxNum = nums[0]
	for _, v := range nums[1:] {
		if v > maxNum {
			maxNum = v
		}
	}
	return
}

func main() {
	l := []int{12, 45, 33, 76, 98, 66, 49}
	fmt.Printf("min: %d\nmax: %d", simpleMin(l), simpleMax(l))
}
