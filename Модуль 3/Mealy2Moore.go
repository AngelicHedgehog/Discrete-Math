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
	fi := make([][]int, n)
	for i = 0; i < n; i++ {
		delta[i] = make([]int, x)
		for j = 0; j < x; j++ {
			_, _ = fmt.Scan(&delta[i][j])
		}
	}
	for i = 0; i < n; i++ {
		fi[i] = make([]int, x)
		for j = 0; j < x; j++ {
			_, _ = fmt.Scan(&fi[i][j])
		}
	}

	fmt.Print("digraph {\n\trankdir = LR")

	done := make([]map[int]bool, n)
	nums := make([]map[int]int, n)
	for i = 0; i < n; i++ {
		done[i] = map[int]bool{}
		nums[i] = map[int]int{}
	}

	for i = 0; i < n; i++ {
		for j = 0; j < x; j++ {
			done[delta[i][j]][fi[i][j]] = true
		}
	}

	k = 0
	for i = 0; i < n; i++ {
		for j = 0; j < y; j++ {
			if done[i][j] {
				fmt.Printf("\n\t%d [label = \"(%d,%s)\"]", k, i, Y[j])
				nums[i][j] = k
				k++
			}
		}
	}

	for i = 0; i < n; i++ {
		for j = 0; j < y; j++ {
			if done[i][j] {
				for k = 0; k < x; k++ {
					fmt.Printf("\n\t%d -> %d [label = \"%s\"]", nums[i][j], nums[delta[i][k]][fi[i][k]], X[k])
				}
			}
		}
	}
	fmt.Print("\n}")
}
