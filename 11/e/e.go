package main

import "fmt"

func main() {
	var s string
	gap := byte('a' - 'A')
	fmt.Scanf("%s", &s)
	for i := 0; i < len(s); i++ {
		if s[i] > 'Z' {
			fmt.Printf("%c", s[i]-gap)
		} else {
			fmt.Printf("%c", s[i]+gap)
		}
	}
	fmt.Println()
}
