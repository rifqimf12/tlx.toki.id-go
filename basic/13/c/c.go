package main

import "fmt"

var a [9]int
var b [10]bool

func isAZigzagPermutation(currentDepth int) bool {
	if (a[currentDepth] < a[currentDepth-1] && a[currentDepth-1] > a[currentDepth-2]) || (a[currentDepth] > a[currentDepth-1] && a[currentDepth-1] < a[currentDepth-2]) {
		return true
	}
	return false
}

func f(x, depth int) {
	if depth == x {
		for i := 0; i < x; i++ {
			fmt.Printf("%d", a[i])
		}
		fmt.Println()
	} else {
		for i := 1; i <= x; i++ {
			if b[i] {
				continue
			}
			a[depth] = i
			if depth >= 2 && !isAZigzagPermutation(depth) {
				continue
			}
			b[i] = true
			f(x, depth+1)
			b[i] = false
		}
	}
}

func printZigzagPermutation(x int) {
	f(x, 0)
}

func main() {
	var n int
	fmt.Scanf("%d", &n)
	printZigzagPermutation(n)
}
