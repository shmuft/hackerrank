package main

import (
	"fmt"
)

func main() {
	var n uint64

	fmt.Scanf("%v", &n)
	var sum int = 0

	for n > 0 {
		if (n & 0x01) == 0 {
			sum++
		}
		// fmt.Println(n & 0x01)
		// fmt.Println(n)
		n = n >> 1
	}
	// fmt.Println(sum)
	fmt.Println(1 << uint(sum))
}

