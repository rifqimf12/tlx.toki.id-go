package main

import (
	"fmt"
)

type point struct {
	x, y int
}

var (
	m, n              int
	b                 [25][25]int
	taken             [25][25]bool
	firstTakingPoints []point
)

func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func taking(b [25][25]int, l, r, c int) (count int) {
	if r < 0 || c < 0 || r == m || c == n || taken[r][c] {
		return
	}
	if l == b[r][c] {
		taken[r][c] = true
		count++
		count += taking(b, l, r+1, c)
		count += taking(b, l, r-1, c)
		count += taking(b, l, r, c+1)
		count += taking(b, l, r, c-1)
	}
	return
}

func collapse(a [25][25]int) [25][25]int {
	for i := 0; i < n; i++ {
		for l := 0; l < m; l++ {
			if taken[l][i] {
				a[l][i] = 0
			}
		}
		for j := m - 1; j > 0; j-- {
			for k := j - 1; k >= 0; k-- {
				if a[j][i] == 0 && a[k][i] > 0 {
					swap(&a[j][i], &a[k][i])
					break
				}
			}
		}
	}
	return a
}

func main() {
	fmt.Scanf("%d%d", &m, &n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Scanf("%d", &b[i][j])
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if !taken[i][j] {
				if t := taking(b, b[i][j], i, j); t > 1 {
					firstTakingPoints = append(firstTakingPoints, point{x: i, y: j})
				}
			}
		}
	}
	var max int
	for _, p := range firstTakingPoints {
		taken = [25][25]bool{}
		fval := taking(b, b[p.x][p.y], p.x, p.y)
		a := collapse(b)
		taken = [25][25]bool{}
		var smax int
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if !taken[i][j] && a[i][j] != 0 {
					if t := taking(a, a[i][j], i, j); t > smax {
						smax = t
					}
				}
			}
		}
		if m := fval*(fval-1) + smax*(smax-1); m > max {
			max = m
		}
	}
	fmt.Println(max)
}
