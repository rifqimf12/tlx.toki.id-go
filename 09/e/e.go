package main

import (
	"fmt"
)

func main() {
	var a, b [100][100]int
	var n, m, p int
	fmt.Scanf("%d %d %d", &n, &m, &p)
	for i := 0; i < n; i++ {
		for ii := 0; ii < m; ii++ {
			fmt.Scanf("%d", &a[i][ii])
		}
	}
	for i := 0; i < m; i++ {
		for ii := 0; ii < p; ii++ {
			fmt.Scanf("%d", &b[i][ii])
		}
	}
	for i := 0; i < n; i++ {
		for ii := 0; ii < p; ii++ {
			sum := 0
			for iii := 0; iii < m; iii++ {
				sum += a[i][iii] * b[iii][ii]
			}
			if ii > 0 {
				fmt.Printf(" ")
			}
			fmt.Printf("%d", sum)
		}
		fmt.Println()
	}
}
