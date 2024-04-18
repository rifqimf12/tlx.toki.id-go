package main

import (
	"fmt"
	"strings"
)

func main() {
	var s1, s2, s3, s4 string
	fmt.Scanf("%s\n%s\n%s\n%s", &s1, &s2, &s3, &s4)
	s1 = strings.Replace(s1, s2, "", 1)
	i := strings.Index(s1, s3) + len(s3)
	fmt.Println(s1[:i] + s4 + s1[i:])
}
