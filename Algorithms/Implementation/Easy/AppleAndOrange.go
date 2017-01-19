package main

import (
	"fmt"
)

func main() {
	var s, t, a, b, m, n, f, i int64
	fmt.Scan(&s, &t)
	fmt.Scan(&a, &b)
	fmt.Scan(&m, &n)
	count := 0
	for i = 0; i < m; i++ {
		fmt.Scan(&f)
		if a+f >= s && a+f <= t {
			count++
		}
	}
	fmt.Println(count)
	count = 0
	for i = 0; i < n; i++ {
		fmt.Scan(&f)
		if b+f >= s && b+f <= t {
			count++
		}
	}
	fmt.Println(count)

}
