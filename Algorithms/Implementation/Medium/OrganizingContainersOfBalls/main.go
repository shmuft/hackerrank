package main

import (
	"fmt"
	"sort"
)

func main() {
	var q, n int
	fmt.Scanf("%d", &q)
	for j := 0; j < q; j++ {
		fmt.Scanf("%d", &n)
		rows := make([]int, n, n)
		columns := make([]int, n, n)
		for i := 0; i < n; i++ {
			rows[i] = 0
			columns[i] = 0
		}
		var temp int
		for i := 0; i < n; i++ {
			for k := 0; k < n; k++ {
				fmt.Scanf("%d", &temp)
				rows[i] += temp
				columns[k] += temp
			}
		}
		sort.Ints(rows)
		sort.Ints(columns)
		var ok bool = true
		for i := 0; i < n; i++ {
			if rows[i] != columns[i] {
				ok = false
				break
			}
		}
		if ok {
			fmt.Println("Possible")
		} else {
			fmt.Println("Impossible")
		}
	}
}
