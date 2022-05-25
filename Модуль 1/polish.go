package main

import "fmt"

func evalPolish() int {
	var op string
	fmt.Scan(&op)
	if '0' <= op[0] && op[0] <= '9' {
		return int(op[0] - '0')
	}
	a, b := evalPolish(), evalPolish()
	switch op[1] {
	case '+':
		return a + b
	case '-':
		return a - b
	default:
		return a * b
	}
}

func main() {
	fmt.Println(evalPolish())
}
