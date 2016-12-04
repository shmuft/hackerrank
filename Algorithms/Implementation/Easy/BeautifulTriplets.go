package main

import (
	"fmt"
)

func main() {
	var i, j, k, n, d int
	fmt.Scanf("%v %v", &n, &d)
	arr := make([]int, n, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%v", &arr[i])
	}

	i = 0
	j = 1
	k = 2
	var sum int = 0

	for ; i < n-2; i++ {
		if j <= i {
			j = i
		}
		for ; (j < n-2) && (arr[j]-arr[i] < d); j++ {
		}
		if arr[j]-arr[i] == d {
			if k <= j {
				k = j + 1
			}
			for ; (k < n-1) && (arr[k]-arr[j] < d); k++ {
			}
			// fmt.Printf("i = %v, j = %v, k = %v\n", i, j, k)
			if arr[k]-arr[j] == d {
				sum++
			}
		}
	}
	fmt.Println(sum)
}

