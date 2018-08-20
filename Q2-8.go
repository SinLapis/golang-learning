package main

import "fmt"

var stack [5]int
var i = 0

func printStack() {
	fmt.Printf("stack: %v\n", stack[:i])
}

func pop() (a int) {
	if i > 0 {
		i--
		a = stack[i]
	} else {
		fmt.Printf("stack is empty\n")
	}
	printStack()
	return
}

func push(a int) {
	if i < 5 {
		stack[i] = a
		i++

	} else {
		fmt.Printf("stack is full\n")
	}
	printStack()
	return
}

func main() {
	push(1)
	pop()
	pop()
	push(2)
	push(3)
	pop()
	push(4)
	push(5)
	pop()
	push(6)
	push(7)
	push(8)
	pop()
	pop()
	push(9)
	push(10)
	pop()
	push(11)
	push(12)
}
