package main

import (
	"fmt"
)

func main() {
	var N, a, b, c, ind int
	_, _ = fmt.Scan(&N)
	sqrN := N * N

	Map := make([]int, sqrN)
	graph := make([][]int, sqrN)
	for i := 0; i < N; i++ {
		for j := 0; j < sqrN; j += N {
			ind = i + j
			_, _ = fmt.Scan(&a)
			Map[ind] = a
			graph[ind] = make([]int, 0, 4)
			if i > 0 {
				graph[ind] = append(graph[ind], ind-1)
			}
			if i < N-1 {
				graph[ind] = append(graph[ind], ind+1)
			}
			if j > 0 {
				graph[ind] = append(graph[ind], ind-N)
			}
			if j < sqrN-N {
				graph[ind] = append(graph[ind], ind+N)
			}
		}
	}
	ways := make([]int, sqrN)
	ways[0] = Map[0]
	mark := make(map[int]bool)
	inQueue := map[int]bool{Map[0]: true}
	queue := map[int]map[int]bool{Map[0]: {0: true}}
	for true {
		if len(queue) == 0 {
			break
		}
		a = -1
		for b = 0; ; b++ {
			if inQueue[b] {
				for c = range queue[b] {
					a = c
					delete(queue[b], c)
					if len(queue[b]) == 0 {
						delete(inQueue, b)
						delete(queue, b)
					}
					break
				}
				if a != -1 {
					break
				}
			}
		}
		mark[a] = true
		for _, ind = range graph[a] {
			if !mark[ind] {
				if ways[ind] > ways[a]+Map[ind] || ways[ind] == 0 {
					ways[ind] = ways[a] + Map[ind]
				}
				if !inQueue[ways[ind]] {
					queue[ways[ind]] = map[int]bool{ind: true}
					inQueue[ways[ind]] = true
				} else {
					queue[ways[ind]][ind] = true
				}
			}
		}
	}
	fmt.Println(ways[sqrN-1])
}
