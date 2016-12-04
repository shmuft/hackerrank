package main

import (
	"fmt"
)

func main() {
	var n, i, jumps int
	fmt.Scanf("%v", &n)
	arr := make([]int, n, n)
	for i = 0; i < n; i++ {
		fmt.Scanf("%v", &arr[i])
	}
	jumps = 0
	for i = 0; i < n-1; {
		if i == n-2 {
			jumps++
			break
		} else {
			if arr[i+2] == 0 {
				i = i + 2
				jumps++
			} else {
				if arr[i+1] == 0 {
					i++
					jumps++
				} else {
					break
				}
			}
		}
	}
	fmt.Println(jumps)
}

