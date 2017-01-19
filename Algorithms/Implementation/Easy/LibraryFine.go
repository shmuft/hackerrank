package main

import (
	"fmt"
)

func main() {
	var d, m, y int
	var d1, m1, y1, count int
	fmt.Scan(&d, &m, &y)
	fmt.Scan(&d1, &m1, &y1)
	count = 0
	if y == y1 {
		if m == m1 {
			if d > d1 {
				count += 15 * (d - d1)
			}
		} else {
			if m > m1 {
				count += 500 * (m - m1)
			}
		}
	} else {
		if y > y1 {
			count += 10000
		}
		//fmt.Println(count)
	}
	fmt.Println(count)
}
