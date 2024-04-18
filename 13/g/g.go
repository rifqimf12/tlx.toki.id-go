package main

import "fmt"

var (
	m, n int
	b    [25][25]int
	bt   [25][25]bool
)

func look(l, r, c int) (count int) {
	if r < 0 || c < 0 || r == m || c == n || bt[r][c] {
		return
	}
	if l == b[r][c] {
		bt[r][c] = true
		count++
		count += look(l, r+1, c)
		count += look(l, r-1, c)
		count += look(l, r, c+1)
		count += look(l, r, c-1)
	}
	return
}

func main() {
	fmt.Scanf("%d%d", &m, &n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Scanf("%d", &b[i][j])
		}
	}
	var r, c int
	fmt.Scanf("%d%d", &r, &c)
	t := look(b[r][c], r, c)
	fmt.Printf("%d\n", t*(t-1))
}
