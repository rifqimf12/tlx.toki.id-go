package main

import "fmt"

func main() {
	var s string
	var upperNext bool
	gap := 'a' - 'A'
	fmt.Scanf("%s", &s)
	for i := 0; i < len(s); i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			fmt.Printf("_%c", s[i]+byte(gap))
		} else if s[i] == '_' {
			upperNext = true
		} else if upperNext {
			fmt.Printf("%c", s[i]-byte(gap))
			upperNext = false
		} else {
			fmt.Printf("%c", s[i])
		}
	}
	fmt.Println()
}
