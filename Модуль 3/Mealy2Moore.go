package main

import "fmt"

func main() {
	var x, y, n, i, j, k int
	_, _ = fmt.Scan(&x)
	X := make([]string, x)
	for i = 0; i < x; i++ {
		_, _ = fmt.Scan(&X[i])
	}
	_, _ = fmt.Scan(&y)
	Y := make([]string, y)
	for i = 0; i < y; i++ {
		_, _ = fmt.Scan(&Y[i])
	}
	_, _ = fmt.Scan(&n)
	delta := make([][]int, n)
	fi := make([][]string, n)
	for i = 0; i < n; i++ {
		delta[i] = make([]int, x)
		for j = 0; j < x; j++ {
			_, _ = fmt.Scan(&delta[i][j])
		}
	}
	for i = 0; i < n; i++ {
		fi[i] = make([]string, x)
		for j = 0; j < x; j++ {
			_, _ = fmt.Scan(&fi[i][j])
		}
	}

	done := make([]map[string]bool, n)
	for i = 0; i < n; i++ {
		done[i] = map[string]bool{}
	}

	fmt.Print("digraph {\n\trankdir = LR")
	for i = 0; i < n; i++ {
		for j = 0; j < x; j++ {
			if !done[delta[i][j]][fi[i][j]] {
				done[delta[i][j]][fi[i][j]] = true
				for k = 0; k < x; k++ {
					fmt.Printf("\n\t(%d,%s) -> (%d,%s) [label = \"%s\"]", delta[i][j], fi[i][j], delta[delta[i][j]][k], fi[delta[i][j]][k], X[k])
				}
			}
		}
	}
	fmt.Print("\n}")
}
