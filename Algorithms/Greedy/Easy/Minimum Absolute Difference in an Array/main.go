package main

import (
	"fmt"
	"sort"
	"math"
)

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func main()  {
	var n int
    fmt.Scanf("%v", &n)
	arr := make([]int, n,n)
	for i:=0; i<n; i++{
		fmt.Scanf("%v", &arr[i])
	}
	sort.Ints(arr)
	minAbs := math.MaxInt32
	for i:=0; i<n-1; i++{
    	if abs(arr[i+1]-arr[i]) < minAbs {
			minAbs = abs(arr[i+1]-arr[i])
		}
	}
	fmt.Println(minAbs)
}