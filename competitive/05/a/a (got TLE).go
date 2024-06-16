package main

import (
	"fmt"
)

var n, m int
var groups [262144]uint64

func getTotalParcel(eggs uint64) (total uint64) {
	for i := 0; i < n; i++ {
		total += eggs / groups[i]
		if total > uint64(m) {
			break
		}
	}
	return
}

func search(x uint64) (ret uint64) {
	var l, r uint64 = 1, (1 << 60) / uint64(n)
	for l <= r {
		mid := (l + r) / 2
		if fx := getTotalParcel(mid); fx >= x {
			ret = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return
}

func main() {
	fmt.Scanf("%d%d", &n, &m)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &groups[i])
	}
	fmt.Println(search(uint64(m+1)) - search(uint64(m)))
}
