package main

import "fmt"

func rec(a string) int {
	len_a := len(a)
	if len_a == 1 || cash[a] {
		return 0
	}
	n, x, sep := 0, 0, 2
	for i := 2; i < len_a-1; i++ {
		switch a[i] {
		case '(':
			n++
		case ')':
			n--
		}
		if n == 0 {
			x += rec(a[sep : i+1])
			sep = i + 1
		}
	}
	cash[a] = true
	return x + 1
}

var cash = make(map[string]bool)

func main() {
	var input string
	fmt.Scan(&input)
	fmt.Println(rec(input))
}
