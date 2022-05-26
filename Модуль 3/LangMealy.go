package main

import "fmt"

func rec(word string, q, M int, nulls map[int]bool, dict *map[string]bool, delta *[][]int, fi *[][]string) {
	if !(*dict)[word] {
		(*dict)[word] = true
		fmt.Printf("%s ", word)
	}
	var next int
	if len(word) < M {
		for i := 0; i < 2; i++ {
			next = (*delta)[q][i]
			if (*fi)[q][i] == "-" {
				if !nulls[next] {
					nulls[next] = true
					rec(word, next, M, nulls, dict, delta, fi)
					delete(nulls, next)
				}
			} else {
				rec(word+(*fi)[q][i], next, M, map[int]bool{}, dict, delta, fi)
			}
		}
	}
}

func main() {
	var n, q0, M, i, j int
	_, _ = fmt.Scan(&n)
	delta := make([][]int, n)
	fi := make([][]string, n)
	for i = 0; i < n; i++ {
		delta[i] = make([]int, 2)
		for j = 0; j < 2; j++ {
			_, _ = fmt.Scan(&delta[i][j])
		}
	}
	for i = 0; i < n; i++ {
		fi[i] = make([]string, 2)
		for j = 0; j < 2; j++ {
			_, _ = fmt.Scan(&fi[i][j])
		}
	}
	_, _ = fmt.Scan(&q0, &M)

	dict := map[string]bool{"": true}
	rec("", q0, M, map[int]bool{}, &dict, &delta, &fi)
}
