package main

import (
	"fmt"
)

func main() {
	var n, c int
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Printf("%d", c)
			c++
			if c == 10 {
				c = 0
			}
		}
		fmt.Println()
	}
}
