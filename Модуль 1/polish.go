package main

import "fmt"

func eval(op rune, a, b int) int {
	switch op {
	case '+':
		return a + b
	case '-':
		return a - b
	default:
		return a * b
	}
}

func main() {
	var (
		op1, op2, op3 rune
		a, b, c, d    int
	)
	fmt.Scanf("(%c %d (%c %d (%c %d %d)))", &op1, &a, &op2, &b, &op3, &c, &d)
	fmt.Println(eval(op1, a, eval(op2, b, eval(op3, c, d))))
}
