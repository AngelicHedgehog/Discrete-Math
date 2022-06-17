package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)
	var N, x, y, xy, a, min int
	_, _ = fmt.Fscan(stdin, &N)

	cords := make([][2]int, N)
	for i := 0; i < N; i++ {
		_, _ = fmt.Fscan(stdin, &x, &y)
		cords[i] = [2]int{x, y}
	}

	dist := 0.0
	edges := map[int]int{}
	for i := 1; i < N; i++ {
		x, y = cords[0][0]-cords[i][0], cords[0][1]-cords[i][1]
		edges[i] = x*x + y*y
	}
	for len(edges) != 0 {
		min = -1
		for a = range edges {
			if min == -1 || edges[min] > edges[a] {
				min = a
			}
		}
		dist += math.Sqrt(float64(edges[min]))
		delete(edges, min)
		for a = range edges {
			x, y = cords[min][0]-cords[a][0], cords[min][1]-cords[a][1]
			xy = x*x + y*y
			if xy < edges[a] {
				edges[a] = xy
			}
		}
	}
	fmt.Printf("%.2f", dist)
}
