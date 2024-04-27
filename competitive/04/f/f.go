package main

import "fmt"

var (
	m, n  int
	b     [25][25]int
	taken [25][25]bool
)

func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func taking(l, r, c int) (count int) {
	if r < 0 || c < 0 || r == m || c == n || taken[r][c] {
		return
	}
	if l == b[r][c] {
		taken[r][c] = true
		count++
		count += taking(l, r+1, c)
		count += taking(l, r-1, c)
		count += taking(l, r, c+1)
		count += taking(l, r, c-1)
	}
	return
}

func collapse() {
	for i := 0; i < n; i++ {
		for l := 0; l < m; l++ {
			if taken[l][i] {
				b[l][i] = 0
			}
		}
		for j := m - 1; j > 0; j-- {
			for k := j - 1; k >= 0; k-- {
				if b[j][i] == 0 && b[k][i] > 0 {
					swap(&b[j][i], &b[k][i])
					break
				}
			}
		}
	}
}

func main() {
	fmt.Scanf("%d%d", &m, &n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Scanf("%d", &b[i][j])
		}
	}
	var max, maxX, maxY int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if !taken[i][j] {
				if t := taking(b[i][j], i, j); t > max {
					max = t
					maxX = i
					maxY = j
				}
			}
		}
	}
	taken = [25][25]bool{}
	taking(b[maxX][maxY], maxX, maxY)
	collapse()
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if j > 0 {
				fmt.Print(" ")
			}
			if b[i][j] == 0 {
				fmt.Print(".")
			} else {
				fmt.Printf("%d", b[i][j])
			}
		}
		fmt.Println()
	}
}
