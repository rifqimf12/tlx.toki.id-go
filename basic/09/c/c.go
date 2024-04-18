package main

import (
	"fmt"
)

func main() {
	var a [1001]int
	var x, n, max int
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &x)
		a[x]++
	}
	for i := 1; i <= 1000; i++ {
		if a[i] >= max {
			max = a[i]
			x = i
		}
	}
	fmt.Printf("%d\n", x)
}
