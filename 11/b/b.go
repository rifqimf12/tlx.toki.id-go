package main

import (
	"fmt"
	"strings"
)

func main() {
	var s, t string
	fmt.Scanf("%s %s", &s, &t)
	for strings.Contains(s, t) {
		s = strings.Replace(s, t, "", 1)
	}
	fmt.Println(s)
}
