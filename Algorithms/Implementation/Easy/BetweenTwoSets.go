package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	a := make([]int, n, n)
	b := make([]int, m, m)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	for i := 0; i < m; i++ {
		fmt.Scan(&b[i])
	}

	sort.Ints(a)
	sort.Ints(b)

	maxA := a[n - 1]
	minB := b[0]
	maxNum := minB / maxA
	count := 0

	for i := 1; i <= maxNum; i++ {
		ok := true
		for j:=0; j<n; j++{
			if (i*maxA) % a[j] !=0 {
				ok = false
				break
			}
		}
		for j:=0; j<m; j++{
			if b[j] % (i*maxA) !=0 {
				ok = false
				break
			}
		}
		if ok {
			count++
		}
	}
	fmt.Println(count)
}
