package main

import (
	"fmt"
)

func main() {
	var a [100][100]int
	var n, m int
	fmt.Scanf("%d %d", &n, &m)
	for i := 0; i < n; i++ {
		for ii := 0; ii < m; ii++ {
			fmt.Scanf("%d", &a[i][ii])
		}
	}
	for i := 0; i < m; i++ {
		for ii := 0; ii < n; ii++ {
			if ii > 0 {
				fmt.Printf(" ")
			}
			fmt.Printf("%d", a[n-ii-1][i])
		}
		fmt.Println()
	}
}
