package main

import "fmt"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func pow(x, y int) int {
	switch y {
	case 0:
		return 1
	case 1:
		return x
	default:
		return x * pow(x, y-1)
	}
}

func closeness(x1, y1, x2, y2, d int) int {
	return pow(abs(x1-x2), d) + pow(abs(y1-y2), d)
}

func main() {
	maxCloseness := 0
	minCloseness := 2 * pow(100, 3)
	var n, d int
	var x, y [1000]int
	fmt.Scanf("%d%d", &n, &d)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d%d", &x[i], &y[i])
		for ii := i - 1; ii >= 0; ii-- {
			c := closeness(x[ii], y[ii], x[i], y[i], d)
			if c < minCloseness {
				minCloseness = c
			}
			if c > maxCloseness {
				maxCloseness = c
			}
		}
	}
	fmt.Println(minCloseness, maxCloseness)
}
