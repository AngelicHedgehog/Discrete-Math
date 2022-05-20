package main

import "fmt"

func find(last, v int, nomer *int, num, up *[]int, mark *map[int]bool, graph [][]int) {
	(*mark)[v] = true
	(*num)[v] = *nomer
	(*up)[v] = *nomer
	*nomer++
	for _, w := range graph[v] {
		if !(*mark)[w] {
			find(v, w, nomer, num, up, mark, graph)
			if (*up)[w] < (*up)[v] {
				(*up)[v] = (*up)[w]
			}
		} else if (*num)[w] < (*up)[v] && w != last {
			(*up)[v] = (*num)[w]
		}
	}
}

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

	var nomer int
	mark := make(map[int]bool)
	num := make([]int, N, N)
	up := make([]int, N, N)
	count := 0
	for i := 0; i < N; i++ {
		if !mark[i] {
			nomer = 0
			find(i, i, &nomer, &num, &up, &mark, graph)
			count--
		}
	}
	for i := 0; i < N; i++ {
		if num[i] == up[i] {
			count++
		}
	}

	fmt.Println(count)
}
