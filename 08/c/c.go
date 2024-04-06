package main

import (
	"fmt"
)

func main() {
	var n, x int
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &x)
		factor := 0 // factor counter besides 1 and itself
		for j := 2; j*j <= x; j++ {
			if x%j == 0 {
				factor++
				if factor == 2 {
					break
				}
			}
		}
		if factor < 2 {
			fmt.Println("YA")
		} else {
			fmt.Println("BUKAN")
		}
	}
}
