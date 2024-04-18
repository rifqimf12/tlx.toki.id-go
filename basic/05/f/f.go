package main

import (
	"fmt"
)

func main() {
	var n float64
	fmt.Scanf("%f", &n)
	if intn := int(n); n-float64(intn) == 0 {
		fmt.Printf("%d %d\n", intn, intn)
	} else if n > 0 {
		fmt.Printf("%d %d\n", intn, intn+1)
	} else {
		fmt.Printf("%d %d\n", intn-1, intn)
	}
}
