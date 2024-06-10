package main

import "fmt"

func main() {
	var a, b, c uint64
	const m uint64 = 1e6
	fmt.Scanf("%d %d", &a, &b)
	c = 1
	reach1m := false
	if a >= m {
		reach1m = true
	}
	a %= m
	for b > 1 {
		if b%2 != 0 {
			c = (c * a) % m
		}
		if a = (a * a); !reach1m && a >= m {
			reach1m = true
		}
		a = a % m
		b /= 2
	}
	if r := (a * c) % m; reach1m {
		fmt.Printf("%06d\n", r)
	} else {
		fmt.Printf("%d\n", r)
	}
}
