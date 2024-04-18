package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	for i := n; i > 0; i-- {
		if n%i == 0 {
			fmt.Printf("%d\n", i)
		}
	}
}
