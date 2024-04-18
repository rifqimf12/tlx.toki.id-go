package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		for j := i; j < n-1; j++ {
			fmt.Print(" ")
		}
		for j := n - i - 1; j < n; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
