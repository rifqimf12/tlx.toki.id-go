package main

import "fmt"

var a [10]int

func printCombination(n, k int) {
	var f func(int)
	f = func(depth int) {
		if depth == k {
			for i := 1; i <= k; i++ {
				if i > 1 {
					fmt.Printf(" ")
				}
				fmt.Printf("%d", a[i])
			}
			fmt.Println()
			return
		} else {
			for i := a[depth] + 1; i <= n; i++ {
				a[depth+1] = i
				f(depth + 1)
			}
		}
	}
	f(0)
}

func main() {
	var n, k int
	fmt.Scanf("%d%d", &n, &k)
	printCombination(n, k)
}
