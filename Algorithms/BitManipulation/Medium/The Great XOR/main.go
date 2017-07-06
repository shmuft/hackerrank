package main

import "fmt"

func main() {
	var q, x, i, sum, n int64
	fmt.Scanf("%d", &q)

	for i = 0; i < q; i++ {
		sum = 0
		fmt.Scanf("%d", &x)
		n = 1
		for x != 0 {
			if (x & 1) == 0 {
				sum += n
			}
			n = n << 1
			x = x >> 1
		}

		fmt.Println(sum)
	}
}
