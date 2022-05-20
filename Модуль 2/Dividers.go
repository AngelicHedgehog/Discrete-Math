package main

import "fmt"

func main() {
	var x int
	_, _ = fmt.Scan(&x)

	n := 0
	graph := append(make([][]int, 0), make([]int, 0))
	dots := []int{x}
	rangs := []int{0}
	rang := 0

	for i := x / 2; i >= 1; i-- {
		if x%i == 0 {
			graph = append(graph, make([]int, 0))
			rang = -1
			for j := n; j >= 0; j-- {
				if dots[j]%i == 0 {
					if rang == -1 {
						rang = rangs[j]
					}
					if rang == rangs[j] {
						graph[j] = append(graph[j], i)
					}
				}
			}
			n++
			dots = append(dots, i)
			rangs = append(rangs, rang+1)
		}
	}
	fmt.Println("graph {")
	for _, n = range dots {
		fmt.Printf("\t%d\n", n)
	}
	for i, a := range graph {
		for _, n = range a {
			fmt.Printf("\t%d -- %d\n", dots[i], n)
		}
	}
	fmt.Println("}")
}
