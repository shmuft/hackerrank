package main

import "fmt"

func main() {
	var n, t int
	fmt.Scanf("%d", &n)
	fmt.Scanf("%d", &t)
	array := make([]int, n, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &array[i])
	}
	var i, j int
	for k := 0; k < t; k++ {
		maxWidth := 4
		fmt.Scanf("%d", &i)
		fmt.Scanf("%d", &j)
		for segment := i; segment <= j; segment++{
			if array[segment] < maxWidth {
				maxWidth = array[segment]
			}
		}
		fmt.Println(maxWidth)
	}
}
