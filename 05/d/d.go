package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	if n > 0 {
		fmt.Println("positif")
	} else if n == 0 {
		fmt.Println("nol")
	} else {
		fmt.Println("negatif")
	}
}
