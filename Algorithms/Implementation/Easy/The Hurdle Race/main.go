package main

import "fmt"

func main() {
	var n, k, h, maxH int
	fmt.Scanf("%d", &n)
	fmt.Scanf("%d", &k)
	maxH = -1
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &h)
		if maxH < h {
			maxH = h
		}
	}
	if maxH - k > 0 {
		fmt.Println(maxH - k)
	} else {
		fmt.Println(0)
	}
}
