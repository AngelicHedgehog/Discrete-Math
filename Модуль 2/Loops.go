package main

import "fmt"

func main() {
	var n, i, j, a, b int
	_, _ = fmt.Scan(&n)

	labels := map[int]int{}
	graph := map[int][][]int{}
	jumps := map[int]int{}
	var (
		label   int
		command string
		operand int
	)
	for i = 0; i < n; i++ {
		graph[i] = [][]int{{}, {}}
	}
	for i = 0; i < n; i++ {
		_, _ = fmt.Scan(&label, &command)
		labels[label] = i
		if command[0] == 'A' || command[0] == 'B' && i != n-1 {
			graph[i][0] = append(graph[i][0], i+1)
			graph[i+1][1] = append(graph[i+1][1], i)
		}
		if command[0] == 'J' || command[0] == 'B' {
			_, _ = fmt.Scan(&operand)
			jumps[i] = operand
		}
	}
	for i = range jumps {
		j = labels[jumps[i]]
		graph[i][0] = append(graph[i][0], j)
		graph[j][1] = append(graph[j][1], i)
	}

	dominators := make([][]bool, n)
	var queue []int
	for i = 0; i < n; i++ {
		dominators[i] = make([]bool, n)
		queue = []int{0}
		for len(queue) != 0 {
			a = queue[0]
			queue = queue[1:]
			dominators[i][a] = true
			if a != i {
				for _, b = range graph[a][0] {
					if !dominators[i][b] {
						queue = append(queue, b)
					}
				}
			}
		}
	}

	count := 0
	for i = 0; i < n; i++ {
		if i == 0 || !dominators[0][i] {
			for _, j = range graph[i][1] {
				if !dominators[i][j] {
					count++
					break
				}
			}
		}
	}
	fmt.Println(count)
}
