package main

import (
	"fmt"
)

func main() {
	var n, m int

	fmt.Scanf("%v %v", &n, &m)
	arr := make([]int, m, m)
	for i := 0; i < m; i++ {
		fmt.Scanf("%v", &arr[i])
	}
	var dist int

	dist = 0
	// var bound bool = true
	if m > 1 {

		var lowest, x int
		for i := 0; i < m-1; i++ {
			lowest = arr[i]
			x = -1
			for j := i + 1; j < m; j++ {
				if arr[j] < lowest {
					lowest = arr[j]
					x = j
				}
			}
			if x != -1 {
				arr[x], arr[i] = arr[i], arr[x]
			}
		}
		// fmt.Println(arr)
		dist = arr[0]
		if dist < n-1-arr[m-1] {
			dist = n - 1 - arr[m-1]
		}
		// fmt.Println(dist)
		for i := 0; i < m-1; i++ {
			if (arr[i+1]-arr[i])/2 > dist {
				dist = (arr[i+1] - arr[i]) / 2
				// fmt.Println(i)
				// fmt.Println(dist)
				// fmt.Println()

				// bound = false
			}
		}

	} else {
		if n == 1 {
			dist = 0
		} else {
			if arr[0] > n-arr[0] {
				dist = arr[0]
			} else {
				if arr[0] == 0 {
					dist = n - 1
				} else {
					dist = n - arr[0]
				}
			}
		}

	}
	// if bound {
	fmt.Println(dist)
	// } else {
	// fmt.Println(dist / 2)
	// }
}

