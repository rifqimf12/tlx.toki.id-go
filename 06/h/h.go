package main

import (
	"fmt"
)

func main() {
	var n, k int

	fmt.Scanf("%d %d", &n, &k)
	for i := 1; i <= n; i++ {
		if i > 1 {
			fmt.Printf(" ")
		}
		if i%k == 0 {
			fmt.Printf("*")
		} else {
			fmt.Printf("%d", i)
		}
	}
	fmt.Println()
}
