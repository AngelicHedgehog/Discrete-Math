package main

/*
	ERROR IN ALGORITHM
*/

import "fmt"

func main() {
	var n, m, q0, a, i, j int
	_, _ = fmt.Scan(&n, &m, &q0)
	delta := make([][]int, n)
	fi := make([][]string, n)
	T := make([]int, n)
	T_ := make([]int, n)
	for i = 0; i < n; i++ {
		T[i] = -1
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

	nextNum := 0
	lenQueue := 0
	queue := []int{q0}
	for lenQueue != -1 {
		a = queue[lenQueue]
		lenQueue--
		if T[a] == -1 {
			T[a] = nextNum
			T_[nextNum] = a
			nextNum++
			for j = 0; j < m; j++ {
				queue = append(queue, delta[a][j])
				lenQueue++
			}
		}
	}

	fmt.Printf("%d\n%d\n%d", n, m, 0)
	for i = 0; i < n; i++ {
		fmt.Println()
		for j = 0; j < m; j++ {
			fmt.Printf("%d ", T[delta[T_[i]][j]])
		}
	}
	for i = 0; i < n; i++ {
		fmt.Println()
		for j = 0; j < m; j++ {
			fmt.Printf("%s ", fi[T_[i]][j])
		}
	}
}
