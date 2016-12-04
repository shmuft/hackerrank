package main

import (
	"fmt"
	"strconv"
)

func reverse(i int) int {
	s := strconv.Itoa(i)
	runes := []rune(s)
	for i, j := 0, len(runes) - 1; i < j; i, j = i + 1, j - 1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	i1, _ := strconv.Atoi(string(runes))
	return i1
}
func main() {
	var i, j, k, l, r, sum int
	fmt.Scan(&i, &j, &k)
	sum = 0
	for m := i; m <= j; m++ {
		l = reverse(m)
		r = m - l
		if r < 0 {
			r *= -1
		}
		//if r < k {
		//	continue
		//}

		if r % k == 0 {
			sum++
		}
	}
	fmt.Println(sum)
}

