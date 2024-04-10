package main

import "fmt"

// return 0 if divisor is not a prime or x is not divisible
func getPowerOfPrimeFactor(x int, divisor int) (power int) {
	for x%divisor == 0 {
		power++
		x /= divisor
	}
	return
}

const separator = " x "

func main() {
	var n int
	fmt.Scanf("%d", &n)
	x := n
	firstPrimeFactor := true
	for i := 2; x > 1; i++ {
		if p := getPowerOfPrimeFactor(x, i); p > 0 {
			for x%i == 0 {
				x /= i
			}
			if !firstPrimeFactor {
				fmt.Print(separator)
			}
			fmt.Printf("%d", i)
			if p > 1 {
				fmt.Printf("^%d", p)
			}
			firstPrimeFactor = false
		}
	}
	fmt.Println()
}
