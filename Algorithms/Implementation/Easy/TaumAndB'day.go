package main

import (
	"fmt"
)

func main() {
	var t int
	var b, w, x, y, z, lCost, rCost int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		fmt.Scan(&b, &w, &x, &y, &z)
		if x+z < y {
			rCost = x + z
		} else {
			rCost = y
		}

		if y+z < x {
			lCost = y+z
		} else {
			lCost = x
		}

		fmt.Println(lCost*b + rCost*w)
	}
}
