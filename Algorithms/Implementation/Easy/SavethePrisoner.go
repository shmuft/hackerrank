package main

import (
	"fmt"
)

func main() {
	var i, val, n, m, s, t int64
	fmt.Scanf("%v", &t)
	for i = 0; i < t; i++ {
		fmt.Scanf("%v %v %v", &n, &m, &s)

		val = (m+s)%n - 1
		if val == 0 {
			val = n
		}
		fmt.Println(val)
	}
}

