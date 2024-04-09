package main

import "fmt"

func reverse(x int) int {
	temp := x
	ret := 0
	for temp > 0 {
		ret = (ret * 10) + (temp % 10)
		temp = temp / 10
	}

	return ret
}

func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)
	fmt.Printf("%d\n", reverse(reverse(a)+reverse(b)))
}
