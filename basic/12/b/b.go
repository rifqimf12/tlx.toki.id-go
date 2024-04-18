package main

import "fmt"

func f(x int) int {
	if x == 1 {
		return 1
	} else if x%2 == 0 {
		return x / 2 * f(x-1)
	} else {
		return x * f(x-1)
	}
}

func main() {
	var n int
	fmt.Scanf("%d", &n)
	fmt.Printf("%d\n", f(n))
}
