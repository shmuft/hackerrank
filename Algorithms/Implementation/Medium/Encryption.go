package main

import (
	"fmt"
	"math"
)

func main() {
	var s string
	var l1, l2 int
	fmt.Scan(&s)
	l1 = int(math.Sqrt(float64(len(s))))
	l2 = l1

	if l1*l2 < len(s) {
		l2 = l1 + 1
	}
	if l1*l2 < len(s) {
		l1++
	}
	for i := 0; i < l2; i++ {
		for j := 0; j < l1; j++ {
			if i+j*l2 < len(s) {
				fmt.Printf("%v", string(s[i + j*l2]))
			}
		}
		fmt.Print(" ")
	}
}
