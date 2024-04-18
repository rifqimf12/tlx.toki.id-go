package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	if n > 0 && n%2 == 0 {
		fmt.Printf("%d\n", n)
	}
}
