package main

import (
	"fmt"
)

func main() {
	var x [101]int
	var i int
	for i = 0; ; i++ {
		if n, _ := fmt.Scanf("%d", &x[i]); n == 0 {
			break
		}
	}
	for i -= 1; i >= 0; i-- {
		fmt.Printf("%d\n", x[i])
	}
}
