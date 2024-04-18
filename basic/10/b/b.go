package main

import "fmt"

func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func readNNumbers(n int) (nNumbers [1001]int) {
	for i := 1; i <= n; i++ {
		fmt.Scanf("%d", &nNumbers[i])
	}
	return
}

func printNNumbers(nNumbers [1001]int, n int) {
	for i := 1; i <= n; i++ {
		if i > 1 {
			fmt.Print(" ")
		}
		fmt.Printf("%d", nNumbers[i])
	}
	fmt.Println()
}

func main() {
	var n, t int
	var x [2][1001]int
	fmt.Scanf("%d", &n)
	x[0] = readNNumbers(n)
	x[1] = readNNumbers(n)
	fmt.Scanf("%d", &t)
	for i := 0; i < t; i++ {
		var q, qq, c, cc int
		fmt.Scanf("%c %d %c %d", &c, &q, &cc, &qq)
		swap(&x[c-'A'][q], &x[cc-'A'][qq])
	}
	printNNumbers(x[0], n)
	printNNumbers(x[1], n)
}
