package main

import "fmt"

func main() {
	var n, m, q0, i, j, k int
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

	groups := make([][]int, 0)
	idToGroup := make([]int, n)
	countGroups := 0
	for i = 0; i < n; i++ {
		for j = 0; j < countGroups; j++ {
			for k = 0; k < m; k++ {
				if fi[i][k] != fi[groups[j][0]][k] {
					break
				}
			}
			if k == m {
				groups[j] = append(groups[j], i)
				idToGroup[i] = j
				break
			}
		}
		if j == countGroups {
			groups = append(groups, []int{i})
			idToGroup[i] = countGroups
			countGroups++
		}
	}

	var (
		countNewGroups, countSubGroups, groupElement int
		newIdToGroup, group                          []int
		newGroups, subGroups                         [][]int
	)

	for true {
		newIdToGroup = make([]int, n)
		newGroups = make([][]int, 0)
		countNewGroups = 0
		for _, group = range groups {
			subGroups = make([][]int, 0)
			countSubGroups = 0
			for _, groupElement = range group {
				for j = 0; j < countSubGroups; j++ {
					for k = 0; k < m; k++ {
						if idToGroup[delta[groupElement][k]] != idToGroup[delta[subGroups[j][0]][k]] {
							break
						}
					}
					if k == m {
						subGroups[j] = append(subGroups[j], groupElement)
						newIdToGroup[groupElement] = countNewGroups + j
						break
					}
				}
				if j == countSubGroups {
					subGroups = append(subGroups, []int{groupElement})
					newIdToGroup[groupElement] = countNewGroups + countSubGroups
					countSubGroups++
				}
			}
			newGroups = append(newGroups, subGroups...)
			countNewGroups += countSubGroups
		}
		if countGroups == countNewGroups {
			break
		}
		idToGroup = newIdToGroup
		groups = newGroups
		countGroups = countNewGroups
	}

	T := make([]int, countGroups)
	T_ := make([]int, countGroups)
	j = 0
	mark := map[int]bool{}
	queue := []int{q0}
	lastInQueue := 0
	for lastInQueue != -1 {
		k = queue[lastInQueue]
		queue = queue[:lastInQueue]
		lastInQueue--
		if mark[idToGroup[k]] {
			continue
		}
		mark[idToGroup[k]] = true
		T[j] = idToGroup[k]
		T_[idToGroup[k]] = j
		j++
		for i = m - 1; i >= 0; i-- {
			queue = append(queue, delta[k][i])
			lastInQueue++
		}
	}

	alphas := make([]int, m)
	for j = 0; j < m; j++ {
		alphas[j] = 'a' + j
	}

	fmt.Print("digraph {\n\trankdir = LR")
	for i = 0; i < countGroups; i++ {
		for j = 0; j < m; j++ {
			fmt.Printf("\n\t%d -> %d [label = \"%c(%s)\"]", i, T_[idToGroup[delta[groups[T[i]][0]][j]]], alphas[j], fi[groups[T[i]][0]][j])
		}
	}
	fmt.Print("\n}")
}
