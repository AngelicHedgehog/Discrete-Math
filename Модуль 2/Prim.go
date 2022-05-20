package main

import "fmt"

func main() {
	var N, M, a, b, c int
	_, _ = fmt.Scan(&N, &M)

	graph := make([]map[int]int, 0)
	for i := 0; i < N; i++ {
		graph = append(graph, make(map[int]int))
	}

	for i := 0; i < M; i++ {
		_, _ = fmt.Scan(&a, &b, &c)
		graph[a][b] = c
		graph[b][a] = c
	}

	taken := make(map[int]bool)
	indexes := make([]int, 0)
	queue := make([]int, 0)
	taken[0] = true
	n := 0
	for i := range graph[0] {
		indexes = append(indexes, i)
		queue = append(queue, graph[0][i])
		n++
		for j := n - 1; j > 0; j-- {
			if queue[j] < queue[j-1] {
				indexes[j], indexes[j-1] = indexes[j-1], indexes[j]
				queue[j], queue[j-1] = queue[j-1], queue[j]
			} else {
				break
			}
		}
	}
	dist := 0
	var newInd int
	for len(queue) != 0 {
		newInd = indexes[0]
		taken[newInd] = true
		dist += queue[0]
		indexes = indexes[1:]
		queue = queue[1:]
		n--
		for i := n - 1; i >= 0; i-- {
			if indexes[i] == newInd {
				indexes = append(indexes[:i], indexes[i+1:]...)
				queue = append(queue[:i], queue[i+1:]...)
				n--
			}
		}
		for a = range graph[newInd] {
			if !taken[a] {
				indexes = append(indexes, a)
				queue = append(queue, graph[newInd][a])
				n++
				for i := n - 1; i > 0; i-- {
					if queue[i] < queue[i-1] {
						indexes[i], indexes[i-1] = indexes[i-1], indexes[i]
						queue[i], queue[i-1] = queue[i-1], queue[i]
					} else {
						break
					}
				}
			}
		}
	}
	fmt.Println(dist)
}
