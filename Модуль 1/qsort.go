package main

import "fmt"

func qsort(n int,
	less func(i, j int) bool,
	swap func(i, j int)) {

	pairs := make([][2]int, 0, n>>1)
	if n > 1 {
		pairs = append(pairs, [2]int{0, n - 1})
	}
	for len(pairs) > 0 {
		sep, end := pairs[0][0], pairs[0][1]
		for sep < end {
			if less(sep, sep+1) {
				swap(end, sep+1)
				end--
			} else {
				swap(sep, sep+1)
				sep++
			}
		}
		x, y := sep-pairs[0][0], pairs[0][1]-sep
		if x > 1 && y > 1 {
			pairs = append(pairs, [2]int{sep + 1, pairs[0][1]})
			pairs[0] = [2]int{pairs[0][0], sep - 1}
		} else if x > 1 {
			pairs[0] = [2]int{pairs[0][0], sep - 1}
		} else if y > 1 {
			pairs[0] = [2]int{sep + 1, pairs[0][1]}
		} else {
			pairs = pairs[1:]
		}
	}
}

func LESS(i, j int) bool {
	return a[i] < a[j]
}

func SWAP(i, j int) {
	a[i], a[j] = a[j], a[i]
}

var a [10]int

func main() {
	a = [...]int{5, 1, 2, 7, 4, 3, 8, 0, 6, 9}
	qsort(len(a), LESS, SWAP)
	for _, x := range a {
		fmt.Printf("%d ", x)
	}
}
