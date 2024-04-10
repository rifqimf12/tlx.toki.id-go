package main

import "fmt"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func f(a, b, x int) int {
	return abs(a*x + b)
}

func main() {
	var a, b, k, x int
	fmt.Scanf("%d%d%d%d", &a, &b, &k, &x)
	for ; k > 0; k-- {
		x = f(a, b, x)
	}
	fmt.Println(x)
}
