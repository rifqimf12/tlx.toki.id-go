package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	for i := 1; i <= n; i++ {
		if i%10 == 0 {
			continue
		} else if i == 42 {
			fmt.Printf("ERROR\n")
			break
		}
		fmt.Printf("%d\n", i)
	}
}
