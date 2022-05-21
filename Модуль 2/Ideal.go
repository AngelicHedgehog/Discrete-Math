package main

import "fmt"

func compare(a, b []int) bool {
	if len(a) == 0 || len(a) > len(b) {
		return true
	}
	if len(a) < len(b) {
		return false
	}
	for i := range a {
		if a[i] > b[i] {
			return true
		}
		if a[i] < b[i] {
			return false
		}
	}
	return false
}

func main() {
	var n, m, a, b, c int
	_, _ = fmt.Scan(&n, &m)

	graph := make([]map[int]int, n)
	for i := 0; i < n; i++ {
		graph[i] = make(map[int]int)
	}
	for i := 0; i < m; i++ {
		_, _ = fmt.Scan(&a, &b, &c)
		a--
		b--
		if (graph[a][b] == 0 || graph[a][b] > c) && a != b {
			graph[a][b] = c
			graph[b][a] = c
		}
	}

	var next []int
	mark := map[int]bool{}
	ways := make([][]int, n)
	for i := 0; i < n; i++ {
		ways[i] = []int{}
	}
	inQueue := map[int]bool{0: true}
	queue := map[int]map[int]bool{0: {0: true}}
	for true {
		//fmt.Println(inQueue)
		if len(inQueue) == 0 {
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
		for b = range graph[a] {
			if !mark[b] {
				next = append(ways[a], graph[a][b])
				if compare(ways[b], next) {
					ways[b] = next
					if inQueue[len(ways[b])] {
						queue[len(ways[b])][b] = true
					} else {
						inQueue[len(ways[b])] = true
						queue[len(ways[b])] = map[int]bool{b: true}
					}
				}
			}
		}
	}
	fmt.Println(len(ways[n-1]))
	for _, a = range ways[n-1] {
		fmt.Printf("%d ", a)
	}
}
