package main

import (
	"bufio"
	"fmt"
	"os"
)

func readNewRune(reader *bufio.Reader) rune {
	run := ' '
	for run == ' ' {
		run, _, _ = reader.ReadRune()
	}
	return run
}

func evalPolish(reader *bufio.Reader) int {
	op := readNewRune(reader)
	if '0' <= op && op <= '9' {
		return int(op - '0')
	}
	op = readNewRune(reader)
	a, b := evalPolish(reader), evalPolish(reader)
	readNewRune(reader)
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
	fmt.Println(evalPolish(bufio.NewReader(os.Stdin)))
}
