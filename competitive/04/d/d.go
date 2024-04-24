package main

import (
	"fmt"
)

type move struct {
	mtype int
	mval  int
}

var moves []move
var moveSequence [11]move
var moveUsed [11]bool
var n, r, y, j int
var val [1e5 + 2]int

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func f(depth int) {
	if depth != r {
		for i := 0; i < n; i++ {
			if moveUsed[i] {
				continue
			}
			moveSequence[depth] = moves[i]
			moveUsed[i] = true
			f(depth + 1)
			moveUsed[i] = false
		}
	} else {
		var score int
		var moveYUsed bool
		for i := 0; i < r; i++ {
			if i != 0 {
				switch moveSequence[i-1].mtype {
				case 'B':
					score += moveSequence[i].mval
				case 'P':
					score += (moveSequence[i].mval * 2)
				case 'L':
					score += (moveSequence[i].mval / 2)
				default:
					score += moveSequence[i].mval
					moveYUsed = true
				}
				if moveYUsed {
					score += y
				}
			} else {
				score += moveSequence[i].mval
			}
		}
		val[min(1e5+1, score)]++
	}
}

func main() {
	fmt.Scanf("Subsoal %d\n", &n)
	fmt.Scanf("%d %d %d %d", &n, &r, &y, &j)
	for i := 0; i < n; i++ {
		var m move
		fmt.Scanf("%d %c\n", &m.mval, &m.mtype)
		moves = append(moves, m)
	}
	f(0)
	for i := int(1e5); i > 0; i-- {
		val[i] += val[i+1]
	}
	for i := 0; i < j; i++ {
		fmt.Scanf("%d", &n)
		fmt.Println(val[n+1])
	}
}
