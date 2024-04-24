package main

import (
	"fmt"
)

var n int
var scores []int
var t [7][7]int

func g(x int) int {
	switch x {
	case 3:
		return 0
	case 1:
		return 1
	default:
		return 3
	}
}

func scoreMatch(i int) bool {
	var score int
	for j := 0; j < n; j++ {
		if i == j {
			continue
		}
		score += t[i][j]
	}
	return score == scores[i]
}

func f(a, b, result int) bool {
	if a > 0 {
		for i := 0; i < a; i++ {
			if !scoreMatch(i) {
				return false
			}
		}
		if a >= n {
			return true
		}
	}
	if b >= n {
		return f(a+1, a+2, 3) || f(a+1, a+2, 1) || f(a+1, a+2, 0)
	}
	t[a][b] = result
	t[b][a] = g(result)
	return f(a, b+1, 3) || f(a, b+1, 1) || f(a, b+1, 0)
}

func main() {
	// now := time.Now()
	// println(now.String())
	var tc int
	fmt.Scanf("%d", &tc)
	for ; tc > 0; tc-- {
		fmt.Scanf("%d", &n)
		for i := 0; i < n; i++ {
			var score int
			fmt.Scanf("%d", &score)
			scores = append(scores, score)
		}
		if f(0, 0, 0) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
		scores = []int{}
	}
	// println(time.Now().String())
	// println(time.Since(now).Seconds())
}
