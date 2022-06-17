package main

import "fmt"

func main() {
	var n, m, q0, i, j int
	_, _ = fmt.Scan(&n, &m, &q0)
	delta := make([][]int, n)
	fi := make([][]string, n)
	for i = 0; i < n; i++ {
		delta[i] = make([]int, m)
		for j = 0; j < m; j++ {
			_, _ = fmt.Scan(&delta[i][j])
		}
	}
	for i = 0; i < n; i++ {
		fi[i] = make([]string, m)
		for j = 0; j < m; j++ {
			_, _ = fmt.Scan(&fi[i][j])
		}
	}
	alphas := make([]int, m)
	for j = 0; j < m; j++ {
		alphas[j] = 'a' + j
	}

	fmt.Print("digraph {\n\trankdir = LR")
	for i = 0; i < n; i++ {
		for j = 0; j < m; j++ {
			fmt.Printf("\n\t%d -> %d [label = \"%c(%s)\"]", i, delta[i][j], alphas[j], fi[i][j])
		}
	}
	fmt.Print("\n}")
}
