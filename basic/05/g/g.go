package main

import (
	"fmt"
)

func main() {
	var x1, y1, x2, y2 int
	fmt.Scanf("%d %d %d %d", &x1, &y1, &x2, &y2)
	fmt.Printf("%d\n", abs(x1-x2)+abs(y1-y2))
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
