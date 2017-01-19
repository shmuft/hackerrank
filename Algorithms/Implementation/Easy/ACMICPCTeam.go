package main

import (
	"fmt"
)

func main() {
	var n, m int
	var s string
	fmt.Scan(&n, &m)
	arrBin := make([][]bool, n, n)
	arrCompare := make(map[int]int)
	for i := 0; i < n; i++ {
		arrBin[i] = make([]bool, m, m)
	}
	for i := 0; i < n; i++ {
		fmt.Scan(&s)
		for ind, val := range s {
			if val == rune('0') {
				arrBin[i][ind] = false
			} else {
				arrBin[i][ind] = true
			}
		}
	}
	var count int
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			count = 0
			for k := 0; k < m; k++ {
				if arrBin[i][k] || arrBin[j][k] {
					count++
				}
			}
			_, ok := arrCompare[count]
			if ok {
				arrCompare[count]++
			} else {
				arrCompare[count] = 1
			}
		}
	}
	maxVal := 0
	maxKey := 0
	for key, val := range arrCompare {
		if key > maxKey {
			maxKey, maxVal = key, val
		}
		if key == maxKey {
			if val > maxVal {
				maxKey, maxVal = key, val
			}
		}
	}
	fmt.Println(maxKey)
	fmt.Println(maxVal)

}
