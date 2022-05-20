package main

import "fmt"

func main() {
	var N, M, a, b int
	_, _ = fmt.Scan(&N, &M)

	graph := make([][2][]int, N) //[i][0] из i-той вершины есть дуга в эти вершины
	for i := 0; i < N; i++ {
		graph[i] = [2][]int{make([]int, 0, N), make([]int, 0, N)}
	}
	for i := 0; i < M; i++ {
		_, _ = fmt.Scan(&a, &b)
		graph[a][0] = append(graph[a][0], b)
		graph[b][1] = append(graph[b][1], a)
	}

	var (
		A, B  map[int]bool
		queue []int
		flag  bool
	)
	treeId := make(map[int]bool)
	for i := 0; i < N; i++ {
		if !treeId[i] {
			A, B = make(map[int]bool), make(map[int]bool)
			queue = append(make([]int, 0), i)
			A[i] = true
			for len(queue) != 0 {
				a = queue[0]
				queue = queue[1:]
				for _, b = range graph[a][0] {
					if !A[b] {
						A[b] = true
						queue = append(queue, b)
					}
				}
			}
			queue = append(queue, i)
			B[i], flag = true, false
			for len(queue) != 0 {
				a = queue[0]
				queue = queue[1:]
				for _, b = range graph[a][1] {
					if !B[b] {
						if A[b] {
							B[b] = true
							queue = append(queue, b)
						} else {
							flag = true
							break
						}
					}
				}
				if flag {
					break
				}
			}
			if !flag {
				a = -1
				for b = range B {
					if b < a || a == -1 {
						a = b
					}
					treeId[b] = true
				}
				fmt.Printf("%d ", a)
			}
		}
	}
}
