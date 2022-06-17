package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)
	var n, m, q0, a, i, j int
	_, _ = fmt.Fscan(stdin, &n, &m, &q0)
	delta := make([][]int, n)
	fi := make([][]string, n)
	T := make([]int, n)
	T_ := make([]int, n)
	for i = 0; i < n; i++ {
		T[i] = -1
		delta[i] = make([]int, m)
		for j = 0; j < m; j++ {
			_, _ = fmt.Fscan(stdin, &delta[i][j])
		}
	}
	for i = 0; i < n; i++ {
		fi[i] = make([]string, m)
		for j = 0; j < m; j++ {
			_, _ = fmt.Fscan(stdin, &fi[i][j])
		}
	}

	num := 0
	queue := []int{q0}
	lastElem := 0
	for lastElem != -1 {
		a = queue[lastElem]
		queue = queue[:lastElem]
		lastElem--
		if T[a] == -1 {
			T[a] = num
			T_[num] = a
			num++
			for i = m - 1; i >= 0; i-- {
				if T[delta[a][i]] == -1 {
					queue = append(queue, delta[a][i])
					lastElem++
				}
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
