package main

import (
	"fmt"
)

func main() {
	var n, temp int
	arr := make(map[int]byte)
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&temp)
		arr[temp]++
	}

	for key, val := range arr {
		if val == 1 {
			fmt.Println(key)
		}
	}
}
