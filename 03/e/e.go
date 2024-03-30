package main

import "fmt"

func main() {
	var a, t int
	fmt.Scanf("%d %d", &a, &t)
	fmt.Printf("%.2f", float32(a*t)/2)
}
