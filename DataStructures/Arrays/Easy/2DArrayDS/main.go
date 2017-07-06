package main

import (
	"fmt"
	"math"
)

func main() {
	var n int = 6
	array := make([][]int, n, n)
	for i := 0; i < n; i++ {
		array[i] = make([]int, n, n)
		for j := 0; j < n; j++ {
			fmt.Scanf("%d", &array[i][j])
		}
	}
	var max int = math.MinInt32
	var sum int
	for i := 1; i < 5; i++ {
		for j := 1; j < 5; j++ {
			sum = 0
			sum += array[i-1][j-1]
			sum += array[i-1][j]
			sum += array[i-1][j+1]
			sum += array[i][j]
			sum += array[i+1][j-1]
			sum += array[i+1][j]
			sum += array[i+1][j+1]
			if sum > max {
				max = sum
			}
		}
	}
	fmt.Println(max)
}
