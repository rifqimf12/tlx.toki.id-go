package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	for {
		if n == 1 {
			fmt.Printf("ya\n")
			break
		} else if n%2 == 1 {
			fmt.Printf("bukan\n")
			break
		}
		n /= 2
	}
}
