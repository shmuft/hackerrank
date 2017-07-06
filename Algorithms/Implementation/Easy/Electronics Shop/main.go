package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, m, s int
	fmt.Scanf("%d", &s)
	fmt.Scanf("%d", &n)
	fmt.Scanf("%d", &m)
	keyboards := make([]int, n, n)
	drives := make([]int, m, m)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &keyboards[i])
	}
	for i := 0; i < m; i++ {
		fmt.Scanf("%d", &drives[i])
	}
	sort.Ints(keyboards)
	sort.Ints(drives)
	var maxSum int
	maxSum = -1
	//var ok bool

	for i := 0; i < n; i++ {
		//ok = true
		for j := m - 1; j >= 0; j-- {
			if keyboards[i]+drives[j] <= s {
				if maxSum < keyboards[i]+drives[j] {
					maxSum = keyboards[i] + drives[j]
			//		ok = false
			//		break
				}
			}
		}
		//if ok {
		//	break
		//}
	}
	fmt.Println(maxSum)
}
