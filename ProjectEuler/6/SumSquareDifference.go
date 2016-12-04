package main

import (
	"fmt"
)

func findDifference(number int64) int64 {
	var sum, diff int64
	sum = number
	for i:=int64(number); i>1; i-- {
		diff += sum * (i-1)
		sum += i-1
	}
	diff = diff * 2
	return diff
}

func main() {
	var i, n, val int64
	fmt.Scanf("%v", &n)
	for i=0; i<n; i++ {
		fmt.Scanf("%v", &val)
		diff := findDifference(val)
		fmt.Println(diff)
	}
}
