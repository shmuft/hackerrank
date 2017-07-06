package main

import (
	"fmt"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	var q, x, y, z int
	fmt.Scanf("%d", &q)
	for i := 0; i < q; i++ {
		fmt.Scanf("%d", &x)
		fmt.Scanf("%d", &y)
		fmt.Scanf("%d", &z)
		if abs(x-z) < abs(y-z) {
			fmt.Println("Cat A")
		} else if abs(x-z) > abs(y-z) {
			fmt.Println("Cat B")
		} else {
			fmt.Println("Mouse C")
		}
	}
}
