package main

import "fmt"

func bubble(nums []int) (results []int) {
	var temp int
	swFlag := false
	for i := 9; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if nums[j] > nums[j+1] {
				temp = nums[j]
				nums[j] = nums[j+1]
				nums[j+1] = temp
				swFlag = true
			}
		}
		fmt.Printf("%v\n", nums)
		if swFlag {
			swFlag = false
		} else {
			break
		}
	}
	return nums
}

func main() {
	nums := []int{3, 6, 31, 9, 14, 78, 22, 16, 99, 1}
	fmt.Printf("result: %v", bubble(nums))
}
