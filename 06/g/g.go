package main

import (
	"fmt"
)

func main() {
	var (
		n, x int
		min  = 100000
		max  = -100000
	)
	fmt.Scanf("%d", &n)
	for ; n > 0; n-- {
		fmt.Scanf("%d", &x)
		if x > max {
			max = x
		}
		if x < min {
			min = x
		}
	}
	fmt.Printf("%d %d\n", max, min)
}
