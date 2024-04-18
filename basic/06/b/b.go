package main

import (
	"fmt"
)

func main() {
	var n, x, total int
	fmt.Scanf("%d", &n)
	for ; n != 0; n-- {
		fmt.Scanf("%d", &x)
		total += x
	}
	fmt.Printf("%d\n", total)
}
