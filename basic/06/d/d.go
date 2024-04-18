package main

import (
	"fmt"
)

func main() {
	var x, total int
	for {
		if n, _ := fmt.Scanf("%d", &x); n == 0 {
			break
		}
		total += x
	}
	fmt.Printf("%d\n", total)
}
