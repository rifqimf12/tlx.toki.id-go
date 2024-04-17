package main

import "fmt"

var a [128][128]int

func isHomogen(r, c, l int) bool {
	for i := r; i < r+l; i++ {
		for j := c; j < c+l; j++ {
			if a[i][j] == 0 {
				return false
			}
		}
	}
	return true
}

func getQuadTreeCode(loc string, r, c, l int) (codes []string) {
	if isHomogen(r, c, l) {
		codes = append(codes, "1"+loc)
	} else if l > 1 {
		l /= 2
		codes = append(codes, getQuadTreeCode(loc+"0", r, c, l)...)
		codes = append(codes, getQuadTreeCode(loc+"1", r, c+l, l)...)
		codes = append(codes, getQuadTreeCode(loc+"2", r+l, c, l)...)
		codes = append(codes, getQuadTreeCode(loc+"3", r+l, c+l, l)...)
	}
	return
}

func getNearestSquare(r, c int) int {
	if r < c {
		r = c
	}
	for i := 2; i <= 128; i = i * 2 {
		if r <= i {
			return i
		}
	}
	return r
}

func main() {
	var r, c int
	fmt.Scanf("%d%d", &r, &c)
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			fmt.Scanf("%d", &a[i][j])
		}
	}
	p := getNearestSquare(r, c)
	q := getQuadTreeCode("", 0, 0, p)
	fmt.Println(len(q))
	for _, s := range q {
		fmt.Println(s)
	}
}
