package main

import (
	"fmt"
	"math"
)

func main() {
	var t, a, b, i, a1, b1 int
	fmt.Scan(&t)
	for i = 0; i < t; i++ {
		fmt.Scan(&a, &b)
		a1 = int(math.Sqrt(float64(a)))
		b1 = int(math.Sqrt(float64(b)))
		if a1*a1 < a && a1*a1 <= b {
			a1++
		}

		count := b1 - a1 + 1
		//fmt.Printf("a1 = %v, b1 = %v\n", a1, b1)

		fmt.Println(count)
	}
}
