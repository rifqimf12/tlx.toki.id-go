package main

import "fmt"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func f(a, b, k, x int) int {
	if k == 0 {
		return abs(a*x + b)
	} else {
		return abs(a*f(a, b, k-1, x) + b)
	}
}

func main() {
	var a, b, k, x int
	fmt.Scanf("%d%d%d%d", &a, &b, &k, &x)
	fmt.Printf("%d\n", f(a, b, k-1, x))
}
