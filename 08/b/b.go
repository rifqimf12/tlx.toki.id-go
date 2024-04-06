package main

import (
	"fmt"
)

func main() {
	var n, x int
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &x)
		isAPrime := true
		for j := 2; j*j <= x; j++ {
			if x%j == 0 {
				isAPrime = false
				break
			}
		}
		if !isAPrime || x == 1 {
			fmt.Println("BUKAN")
		} else {
			fmt.Println("YA")
		}
	}
}
