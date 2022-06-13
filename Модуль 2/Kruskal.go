package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

/*
	TIME LIMIT
*/

func main() {
	stdin := bufio.NewReader(os.Stdin)
	var N, a, b int
	_, _ = fmt.Fscan(stdin, &N)

	cords := make([][2]int, N)
	for i := 0; i < N; i++ {
		_, _ = fmt.Fscan(stdin, &a, &b)
		cords[i] = [2]int{a, b}
	}

	var (
		edges                    [][2]int
		lenEdges, sortedLenEdges []int
	)
	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			edges = append(edges, [2]int{i, j})
			a = (cords[i][0]-cords[j][0])*(cords[i][0]-cords[j][0]) + (cords[i][1]-cords[j][1])*(cords[i][1]-cords[j][1])
			lenEdges = append(lenEdges, a)
			sortedLenEdges = append(sortedLenEdges, a)
		}
	}
	sort.Ints(sortedLenEdges)

	dist := 0.0
	comps := make([]int, N)
	for i := range comps {
		comps[i] = i
	}
	for _, a = range sortedLenEdges {
		for i := range lenEdges {
			if a == lenEdges[i] {
				if comps[edges[i][0]] != comps[edges[i][1]] {
					dist += math.Sqrt(float64(a))
					for j := 0; j < N; j++ {
						if comps[j] == comps[edges[i][0]] && j != edges[i][0] {
							comps[j] = comps[edges[i][1]]
						}
					}
					comps[edges[i][0]] = comps[edges[i][1]]
				}
				lenEdges[i] = -1
			}
		}
	}

	fmt.Printf("%.2f", dist)
}
