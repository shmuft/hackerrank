package main

import (
	"fmt"
)

func main() {
	var t, n, r int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		fmt.Scan(&n)
		r = n
		count := 0
		for r > 0 {
			if (r % 10 != 0 && n % ( r % 10 ) == 0) {
				count++
			}
			r = r / 10
		}
		fmt.Println(count)
	}
}
