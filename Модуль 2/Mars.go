package main

import (
	"fmt"
	"sort"
)

/*
	ERROR IN ALGORITHM
*/

func checkByWidth(dot, n int, graph map[[2]int]bool) (bool, [2]map[int]bool) {
	var (
		i, now, team int
		index        [2]int
	)
	queue := append(make([]int, 0), dot)
	res := [2]map[int]bool{make(map[int]bool), make(map[int]bool)}
	res[0][dot] = true
	for len(queue) != 0 {
		now = queue[0]
		queue = queue[1:]
		index[0] = now
		for i = 1; i <= n; i++ {
			index[1] = i
			if graph[index] {
				if res[0][now] {
					team = 1
				} else {
					team = 0
				}
				if res[(team+1)%2][i] {
					return true, [2]map[int]bool{}
				}
				if !res[team][i] {
					queue = append(queue, i)
				}
				res[team][i] = true
			}
		}
	}
	return false, res
}

func add(team1, team2 *map[int]bool, a, b map[int]bool) {
	for x := range a {
		(*team1)[x] = true
	}
	for y := range b {
		(*team2)[y] = true
	}
}

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	var c string
	graph := make(map[[2]int]bool)
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if _, _ = fmt.Scan(&c); c == "+" {
				graph[[2]int{i, j}] = true
			}
		}
	}

	team1, team2 := make(map[int]bool), make(map[int]bool)
	counts := make([]int, 0)
	sumCounts := 0
	subTeams := make([][2]map[int]bool, 0)
	iter := make(map[int]bool)
	var (
		err                   bool
		out                   [2]map[int]bool
		lenA, lenB, diff, sum int
	)
	for i := 1; i <= n; i++ {
		if !iter[i] {
			err, out = checkByWidth(i, n, graph)
			if err {
				fmt.Println("No solution")
				return
			}
			for _, x := range out {
				for y := range x {
					iter[y] = true
				}
			}
			lenA, lenB = len(out[0]), len(out[1])
			if lenA == lenB {
				add(&team1, &team2, out[0], out[1])
			} else {
				if lenA > lenB {
					diff = lenA - lenB
				} else {
					diff = lenB - lenA
					out[0], out[1] = out[1], out[0]
				}
				counts = append(counts, diff)
				sumCounts += diff
				subTeams = append(subTeams, out)
				for j := len(counts) - 1; j > 0; j-- {
					if counts[j] < counts[j-1] {
						counts[j], counts[j-1] = counts[j-1], counts[j]
						subTeams[j], subTeams[j-1] = subTeams[j-1], subTeams[j]
					} else {
						break
					}
				}
			}
		}
	}
	if sumCounts != 0 {
		sum = 0
		for j := len(counts) - 1; j >= 0; j-- {
			sum += counts[j] * 2
			if sum > sumCounts {
				if sum-sumCounts >= sumCounts-sum+counts[j]*2 {
					j++
				}
				a, b := make(map[int]bool), make(map[int]bool)
				for i := len(counts) - 1; i >= 0; i-- {
					if i < j {
						add(&a, &b, subTeams[i][1], subTeams[i][0])
					} else {
						add(&a, &b, subTeams[i][0], subTeams[i][1])
					}
				}
				if len(a) == len(b) {
					for g := 1; g <= n; g++ {
						if a[g] {
							add(&team1, &team2, a, b)
							break
						} else if b[g] {
							add(&team1, &team2, b, a)
							break
						}
					}
				} else if len(a) < len(b) {
					add(&team1, &team2, a, b)
				} else {
					add(&team1, &team2, b, a)
				}
				break
			}
		}
	}

	final := make([]int, 0, len(team1))
	for X := range team1 {
		final = append(final, X)
	}
	sort.Ints(final)
	for _, X := range final {
		fmt.Printf("%d ", X)
	}
}
