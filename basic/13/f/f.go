package main

import "fmt"

var a [128][128]int

func initialize(code string, r, c, l int) {
	if code == "" {
		for i := r; i < r+l; i++ {
			for j := c; j < c+l; j++ {
				a[i][j] = 1
			}
		}
	} else {
		l /= 2
		switch code[0] {
		case '0':
			initialize(code[1:], r, c, l)
		case '1':
			initialize(code[1:], r, c+l, l)
		case '2':
			initialize(code[1:], r+l, c, l)
		case '3':
			initialize(code[1:], r+l, c+l, l)
		}
	}
}

func getNearestSquare(r, c int) int {
	if r < c {
		r = c
	}
	for i := 2; i <= 128; i *= 2 {
		if r <= i {
			return i
		}
	}
	return r
}

func main() {
	var n, r, c int
	var codes []string
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		var s string
		fmt.Scanf("%s\n", &s)
		codes = append(codes, s)
	}
	fmt.Scanf("%d%d", &r, &c)
	l := getNearestSquare(r, c)
	for _, code := range codes {
		initialize(code[1:], 0, 0, l)
	}
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if j > 0 {
				fmt.Printf(" ")
			}
			fmt.Printf("%d", a[i][j])
		}
		fmt.Println()
	}
}
