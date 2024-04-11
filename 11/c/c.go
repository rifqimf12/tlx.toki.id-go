package main

import "fmt"

func main() {
	var s string
	var k int
	fmt.Scanf("%s\n%d", &s, &k)
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c", (s[i]-97+byte(k))%26+97)
	}
	fmt.Println()
}
