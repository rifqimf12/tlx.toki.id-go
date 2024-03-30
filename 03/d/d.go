package main

import "fmt"

func main() {
	var n, m int
	fmt.Scanf("%d %d", &n, &m)
	fmt.Printf("masing-masing %d\nbersisa %d\n", n/m, n%m)
}
