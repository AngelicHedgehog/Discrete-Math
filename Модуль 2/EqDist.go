package main

import "fmt"

func main() {
	var N, M, K, a, b int
	_, _ = fmt.Scan(&N, &M)

	graph := make([][]int, N)
	for i := 0; i < N; i++ {
		graph[i] = make([]int, 0)
	}

	for i := 0; i < M; i++ {
		_, _ = fmt.Scan(&a, &b)
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	_, _ = fmt.Scan(&K)
	sups := make([]int, K)
	for i := 0; i < K; i++ {
		_, _ = fmt.Scan(&sups[i])
	}

	var (
		finDists map[int]int
		dists    map[int]int
		queue    []int
	)
	for i, sup := range sups {
		dists = map[int]int{sup: 1}
		queue = []int{sup}
		for len(queue) != 0 {
			for _, a = range graph[queue[0]] {
				if dists[a] == 0 {
					queue = append(queue, a)
					dists[a] = dists[queue[0]] + 1
				}
			}
			queue = queue[1:]
		}
		if i == 0 {
			finDists = dists
		} else {
			for a = range finDists {
				if finDists[a] != dists[a] {
					delete(finDists, a)
				}
			}
		}
	}
	if len(finDists) == 0 {
		fmt.Println("-")
	} else {
		for i := 0; i < N; i++ {
			if finDists[i] != 0 {
				fmt.Printf("%d ", i)
			}
		}
	}
}
