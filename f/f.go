package main

import "fmt"

func main() {
	var a, b, c, d, e, f, g, h, i int
	fmt.Scanf("%d %d %d\n%d %d %d\n%d %d %d", &a, &b, &c, &d, &e, &f, &g, &h, &i)
	fmt.Printf("%d %d %d\n%d %d %d\n%d %d %d\n", a, d, g, b, e, h, c, f, i)
}
