package main

import "fmt"

func main() {
	var N, M, a, b int
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

	maxDots, maxEdges := -1, -1
	doneDots := make(map[int]bool)
	var (
		dots, edges   int
		maxComp, comp map[int]bool
		queue         []int
	)
	for i := 0; i < N; i++ {
		if !doneDots[i] {
			dots, edges = 1, 0
			comp = map[int]bool{i: true}
			queue = []int{i}
			for len(queue) != 0 {
				a = queue[0]
				queue = queue[1:]
				for _, b = range graph[a] {
					edges++
					if !comp[b] {
						dots++
						comp[b] = true
						doneDots[b] = true
						queue = append(queue, b)
					}
				}
			}
			if dots > maxDots || dots == maxDots && edges > maxEdges {
				maxDots = dots
				maxEdges = edges
				maxComp = comp
			}
		}
	}

	fmt.Println("graph {")
	for i := 0; i < N; i++ {
		if maxComp[i] {
			fmt.Printf("\t%d [color=red]\n", i)
		} else {
			fmt.Printf("\t%d\n", i)
		}
	}
	for i := 0; i < N; i++ {
		for _, a = range graph[i] {
			if a >= i {
				if maxComp[i] {
					fmt.Printf("\t%d -- %d [color=red]\n", i, a)
				} else {
					fmt.Printf("\t%d -- %d\n", i, a)
				}
			}
		}
	}
	fmt.Println("}")
}
