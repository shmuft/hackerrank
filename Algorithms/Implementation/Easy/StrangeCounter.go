package main

import (
	"fmt"
	"math"
)

func main() {
	var t int64
	fmt.Scanf("%v", &t)

	var n, a1 float64

	a1 = 3
	// for t = 1; t < 20; t++ {
	n = math.Log2(float64(t+2) / a1)
	// fmt.Printf("log2(%v) = %v\n", t, n)
	fmt.Printf("%.f", a1*math.Pow(2, float64(int64(n))+1) - float64(t+2))
	// }
}

