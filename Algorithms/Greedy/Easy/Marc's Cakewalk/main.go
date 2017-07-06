package main

import (
	"fmt"
	"sort"
)


func main() {
	var n int
	fmt.Scanf("%d", &n)
	cupCakes := make([]int, n, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &cupCakes[i])
	}
	sort.Ints(cupCakes)
	var sum int = 0

	for i := 0; i < n; i++ {
		sum += cupCakes[n-1-i] * (1 << uint(i))
	}

	fmt.Println(sum)
}
