package main

import (
	"fmt"
)

func main() {
	var t, y int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		h := 0
		fmt.Scan(&y)
		val := true
		for j := 0; j <= y; j++ {
			if val {
				h++
			} else {
				h *= 2
			}
			val = !val
		}
		fmt.Println(h)
	}

}
