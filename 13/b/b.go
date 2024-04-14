package main

import "fmt"

func f(x int) {
	if x > 1 {
		f(x - 1)
	}
	for i := 0; i < x; i++ {
		fmt.Printf("*")
	}
	fmt.Println()
	if x > 1 {
		f(x - 1)
	}
}

func main() {
	var n int
	fmt.Scanf("%d", &n)
	f(n)
}
