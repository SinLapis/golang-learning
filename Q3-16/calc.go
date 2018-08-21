package main

func calculate(f byte) {
	_, b := Pop()
	_, a := Pop()
	var c int
	switch f {
	case '+':
		c = a + b
	case '-':
		c = a - b
	case '*':
		c = a * b
	case '/':
		c = a / b
	}
	Push(c)
}

func main() {
	Push(5)
	Push(1)
	Push(2)
	calculate('+')
	Push(4)
	calculate('*')
	calculate('+')
	Push(3)
	calculate('-')
}
