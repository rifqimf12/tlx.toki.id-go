package main

import "fmt"

func getBinary(x int) string {
	if x == 1 {
		return "1"
	} else if x%2 == 1 {
		return getBinary(x/2) + "1"
	} else {
		return getBinary(x/2) + "0"
	}
}

func main() {
	var n int
	fmt.Scanf("%d", &n)
	fmt.Printf("%s\n", getBinary(n))
}
