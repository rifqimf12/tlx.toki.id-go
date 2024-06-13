package main

import "fmt"

var a, b, c, n int

func f(a, b int) int {
	if b == 0 {
		return 1
	}
	if b == 1 {
		return a
	}
	if b%2 == 1 {
		return modOfMultiply(f(a, b-1), a)
	} else {
		temp := f(a, b/2)
		return modOfMultiply(temp, temp)
	}
}

func modOfMultiply(a, b int) int {
	return int((uint64(a) * uint64(b)) % uint64(n))
}

func main() {
	fmt.Scanf("%d%d%d%d", &a, &b, &c, &n)
	if c == 0 {
		fmt.Printf("%d\n", (a%n)+1)
	} else if c == 1 {
		fmt.Printf("%d\n", (f(a, b)%n)+1)
	} else {
		for i := 0; i < c; i++ {
			a = f(a, b) % n
		}
		fmt.Printf("%d\n", a+1)
	}
}
