package main

import "fmt"

var n, x, q, y int
var z [1e5 + 1]int

func main() {
	fmt.Scanf("%d", &n)
	for i := 1; i <= n; i++ {
		fmt.Scanf("%d", &x)
		z[i] = z[i-1] + x
	}
	fmt.Scanf("%d", &q)
	for i := 0; i < q; i++ {
		fmt.Scanf("%d", &y)
		l, r, ret := 1, n, 0
		for l <= r {
			m := (l + r) / 2
			// println(l, r, m, z[m])
			if y <= z[m] {
				ret = m
				r = m - 1
			} else {
				l = m + 1
			}
		}
		fmt.Println(ret)
	}
}
