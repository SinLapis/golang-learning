package main

import "fmt"

func plus22(a interface{}) (b int) {
	return a.(int) + 2
}

func lenPlus12(str interface{}) (al int) {
	return len(str.(string)) + 1
}

func simpleMap2(f func(interface{}) int, args []interface{}) (results []int) {
	results = make([]int, len(args))
	for k, v := range args {
		results[k] = f(v)
	}
	return
}

func main() {
	l1 := []interface{}{1, 3, 5, 11, 45}
	l2 := []interface{}{"ab", "reedsas", "dddd"}
	fmt.Printf("%v\n", simpleMap2(plus22, l1))
	fmt.Printf("%v\n", simpleMap2(lenPlus12, l2))
}
