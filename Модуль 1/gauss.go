package main

import "fmt"

func gcd(a, b int) int {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	if b > a {
		a, b = b, a
	}
	for a > 0 {
		a, b = b%a, a
	}
	return b
}

func main() {
	/**/
	var N int
	fmt.Scanf("%d", &N)
	slau := make([][]int, N)
	for i := 0; i < N; i++ {
		slau[i] = make([]int, N+1)
	}
	for i := range slau {
		for j := range slau[i] {
			fmt.Scan(&slau[i][j])
		}
	}
	/*/
	N := 5
	slau := [5][6]int{{3, 0, 1, -2, 0, 0}, {-1, 2, -1, 1, -1, -3}, {2, 2, 0, -1, -1, -3}, {1, -1, -1, 0, -2, 1}, {0, 0, 0, 0, 0, 0}}
	/**/

	var d, k1, k2, i, j, x, y int
	for x, y = 0, 0; x < N && y < N; {
		for i = y; i < N; i++ {
			if slau[i][x] != 0 {
				slau[y], slau[i] = slau[i], slau[y]
				i--
				break
			}
		}
		if i == N {
			x++
			continue
		}
		for i = y + 1; i < N; i++ {
			d = gcd(slau[y][x], slau[i][x])
			k1, k2 = slau[y][x]/d, slau[i][x]/d
			for j = x; j <= N; j++ {
				slau[i][j] = slau[i][j]*k1 - slau[y][j]*k2
			}
		}
		x++
		y++
	}

	for y = N - 1; y >= 0; y-- {
		for x = 0; x < N && slau[y][x] == 0; x++ {
		}
		if x == N {
			if slau[y][N] != 0 {
				fmt.Println("No solution")
				return
			}
		} else {
			for i = 0; i < y; i++ {
				d = gcd(slau[y][x], slau[i][x])
				k1, k2 = slau[y][x]/d, slau[i][x]/d
				for j = 0; j <= N; j++ {
					slau[i][j] = slau[i][j]*k1 - slau[y][j]*k2
				}
			}
		}
	}
	for i, j = 0, 0; i < N; {
		if slau[i][j] == 0 {
			if j == N-1 {
				fmt.Println("0/1")
				i++
			} else {
				j++
			}

		} else {
			d = gcd(slau[i][j], slau[i][N])
			x, y = slau[i][N]/d, slau[i][j]/d
			if y < 0 {
				x = -x
				y = -y
			}
			fmt.Printf("%d/%d\n", x, y)
			i++
			j++
		}
	}
	/*
		for i = 0; i < N; i++ {
			for j = 0; j <= N; j++ {
				fmt.Printf("%d ", slau[i][j])
			}
			fmt.Println()
		}
	*/
}
