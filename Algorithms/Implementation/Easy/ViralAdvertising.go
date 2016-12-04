package main

import (
	"fmt"
)

func main() {
	var n, i, sum, num int

	fmt.Scanf("%v", &n)


	sum = 0
	start := 5
	for i = 0; i<n; i++{
		if (float64(start) / 2) - float64(start / 2) > float64(0.5) {
			num = start / 2 + 1
		} else {
			num = start / 2
		}
		sum += num
		start = num * 3
	}
	fmt.Println(sum)

}

