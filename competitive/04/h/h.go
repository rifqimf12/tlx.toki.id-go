package main

import (
	"fmt"
	"sort"
)

var m, n, minMember, nMaxMember int
var b []int
var l, memberCount [20]int

func getGroupedMember(depth int) (t int) {
	for i := 0; i <= depth; i++ {
		t += memberCount[i]
	}
	return
}

func satisfy(depth int) bool {
	var c, cd int
	t := getGroupedMember(depth)
	for i := 0; i < t; i++ {
		if b[i] < l[cd] {
			c++
		} else if c >= minMember && c <= minMember+1 {
			c = 1
			cd++
		} else {
			return false
		}
	}
	if t < len(b) && b[t] < l[depth] {
		return false
	}
	return true
}

func f(curMaxMember, depth int) bool {
	groupedMember := getGroupedMember(depth)
	// max
	if curMaxMember < nMaxMember {
		curMaxMember++
		l[depth] = b[groupedMember+minMember] + 1
		memberCount[depth] = minMember + 1
		if satisfy(depth) {
			if depth < n-1 {
				if f(curMaxMember, depth+1) {
					return true
				}
			} else {
				return true
			}
		}
		memberCount[depth] = 0
		curMaxMember--
	}
	// min
	l[depth] = b[groupedMember+minMember-1] + 1
	memberCount[depth] = minMember
	if satisfy(depth) {
		if depth < n-1 {
			if f(curMaxMember, depth+1) {
				return true
			}
		} else {
			return true
		}
	}
	memberCount[depth] = 0
	return false
}

func main() {
	fmt.Scanf("%d", &m)
	for i := 0; i < m; i++ {
		var x int
		fmt.Scanf("%d", &x)
		b = append(b, x)
	}
	fmt.Scanf("%d", &n)
	minMember = m / n
	nMaxMember = m % n
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})
	f(0, 0)
	for i := 0; i < n-1; i++ {
		fmt.Print(l[i], " ")
	}
	fmt.Println()
}
