package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	switch {
	case n < 10:
		fmt.Println("satuan")
	case n < 100:
		fmt.Println("puluhan")
	case n < 1000:
		fmt.Println("ratusan")
	case n < 10000:
		fmt.Println("ribuan")
	default:
		fmt.Println("puluhribuan")
	}
}
