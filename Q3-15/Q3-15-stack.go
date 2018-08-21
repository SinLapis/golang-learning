package Q3_15

import "fmt"

var Stack [10]int
var i2 = 0

func printStack2() {
	fmt.Printf("stack: %v\n", Stack[:i2])
}

func Pop() (status bool, a int) {
	if i2 > 0 {
		i2--
		a = Stack[i2]
		status = true
	} else {
		fmt.Printf("stack is empty\n")
		a = 0
		status = false
	}
	printStack2()
	return
}

func Push(a int) (status bool) {
	if i2 < 5 {
		Stack[i2] = a
		i2++
		status = true
	} else {
		fmt.Printf("stack is full\n")
		status = false
	}
	printStack2()
	return
}
