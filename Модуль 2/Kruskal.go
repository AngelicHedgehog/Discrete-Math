package main

import (
	"fmt"
	"math"
)

func main() {
	var N, a, b int
	_, _ = fmt.Scan(&N)

	cords := make([][2]int, N, N)
	for i := 0; i < N; i++ {
		_, _ = fmt.Scan(&a, &b)
		cords[i] = [2]int{a, b}
	}

	NN := N * (N - 1) / 2
	dots := make([][2]int, NN, NN)
	edges := make([]int, NN, NN)

	n := 0
	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			dots[n] = [2]int{i, j}
			a, b = cords[i][0]-cords[j][0], cords[i][1]-cords[j][1]
			edges[n] = a*a + b*b
			n++
			for k := n - 1; k > 0; k-- {
				if edges[k] < edges[k-1] {
					dots[k], dots[k-1] = dots[k-1], dots[k]
					edges[k], edges[k-1] = edges[k-1], edges[k]
				}
			}
		}
	}
	treeId := make([]int, 0)
	for i := 0; i < N; i++ {
		treeId = append(treeId, i)
	}
	var oldId, newId int
	dist := 0.0
	for i := 0; i < NN; i++ {
		if treeId[dots[i][0]] != treeId[dots[i][1]] {
			dist += math.Sqrt(float64(edges[i]))
			newId, oldId = treeId[dots[i][0]], treeId[dots[i][1]]
			for j := 0; j < N; j++ {
				if treeId[j] == oldId {
					treeId[j] = newId
				}
			}
		}
	}
	fmt.Printf("%.2f", dist)
}
