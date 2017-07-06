package main

import (
	"fmt"
	"sort"
)

func absInt(x int) int {
	if (x < 0) {
		return -x
	} else {
		return x
	}
}

func main() {
	var n, temp int
	fmt.Scanf("%d", &n)
	arr := make(map[int]int)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &temp)
		_, ok := arr[temp]
		if !ok {
			arr[temp] = 1
		} else {
			arr[temp]++
		}
	}
	arr2 := make([]int, len(arr), len(arr))
	i := 0
	var maxKey int
	maxKey = 0
	for key, val := range arr {
		arr2[i] = key
		i++
		if maxKey < val {
			maxKey = val
		}
	}
	sort.Ints(arr2)
	var max int
	max = 0
	if len(arr2) == 1 {
		max = n
	} else {
		for i := 0; i < len(arr2)-1; i++ {
			if arr2[i+1]-arr2[i] <= 1 {
				if max < arr[arr2[i]]+arr[arr2[i+1]] {
					max = arr[arr2[i]] + arr[arr2[i+1]]
				}
			}
		}
	}
	if max < maxKey {
		max = maxKey
	}
	fmt.Println(max)
}
