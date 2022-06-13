package main

import (
	"fmt"
	"sort"
)

/*
	ERROR IN ALGORITHM
*/

func main() {
	var N, M, K, a, b int
	_, _ = fmt.Scan(&N, &M)

	graph := make([][]int, 0)
	for i := 0; i < N; i++ {
		graph = append(graph, make([]int, 0))
	}

	for i := 0; i < M; i++ {
		_, _ = fmt.Scan(&a, &b)
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	_, _ = fmt.Scan(&K)
	sups := make([]int, 0)
	for i := 0; i < K; i++ {
		_, _ = fmt.Scan(&a)
		sups = append(sups, a)
	}

	exists := make(map[int]bool)
	for i := 0; i < N; i++ {
		exists[i] = true
	}
	dists := make([]int, N, N)
	var (
		dots  []int
		done  map[int]bool
		queue []int
		x     int
	)
	for i := 0; i < K; i++ {
		dots = make([]int, N, N)
		done = make(map[int]bool)
		done[sups[i]] = true
		queue = append(queue, sups[i])
		for len(queue) != 0 {
			x = queue[0]
			if exists[x] && dists[x] != dots[x] {
				delete(exists, x)
			}
			queue = queue[1:]
			for _, b = range graph[x] {
				if !done[b] {
					done[b] = true
					dots[b] = dots[x] + 1
					if i == 0 {
						dists[b] = dots[b]
					}
					queue = append(queue, b)
				}
			}
		}
	}

	if len(exists) == 0 {
		fmt.Printf("-")
	} else {
		queue = make([]int, 0, len(exists))
		for a = range exists {
			queue = append(queue, a)
		}
		sort.Ints(queue)
		for _, a = range queue {
			fmt.Printf("%d ", a)
		}
	}
}
