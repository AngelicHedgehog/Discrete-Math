package main

import "fmt"

func add(a, b []int32, p int) []int32 {
	P := int32(p)
	if len(b) > len(a) {
		a, b = b, a
	}

	c := make([]int32, 0, len(a)+1)
	for _, x := range b {
		c = append(c, x)
	}

	for i, x := range a {
		if len(c) == i {
			c = append(c, x)
		} else {
			c[i] += x
			if c[i] >= P {
				c[i] -= P
				if len(c) == i+1 {
					c = append(c, 1)
				} else {
					c[i+1]++
				}
			}
		}
	}

	return c
}

func main() {
	A := []int32{1, 8, 6, 2, 8, 3, 0, 5, 9, 1}
	B := []int32{5, 1, 1, 3, 7, 2, 3, 3, 1, 8, 9}
	for _, x := range add(A, B, 10) {
		fmt.Printf("%d ", x)
	}
}
