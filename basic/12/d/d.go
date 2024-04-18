package main

import "fmt"

func isAPalindrom(s string) bool {
	if len(s) <= 1 {
		return true
	} else if s[0] == s[len(s)-1] {
		return isAPalindrom(s[1 : len(s)-1])
	} else {
		return false
	}
}

func main() {
	var s string
	fmt.Scanf("%s", &s)
	if isAPalindrom(s) {
		fmt.Println("YA")
	} else {
		fmt.Println("BUKAN")
	}
}
