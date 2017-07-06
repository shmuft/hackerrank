package main

import (
	"fmt"
)

func main() {
	var p, q, i, val, l, r, pow int64
	arr := make([]int64, 11, 11)
	var lenValue int = 0
	fmt.Scan(&p, &q)
	var ok bool
	ok = false
	for i = p; i <= q; i++ {
		val = i * i
		for lenValue = 0; val != 0; lenValue++ {
			arr[lenValue] = val % 10
			val /= 10
		}
		pow = 1
		l = 0
		for j := 0; j <= (lenValue-1)/2; j++ {
			l += arr[j] * pow
			pow *= 10
		}
		//fmt.Println(l)

		pow = 1
		r = 0
		j := 0
		if (lenValue % 2) == 0 {
			j = lenValue / 2
		} else {
			j = lenValue/2 + 1
		}
		for ; j <= lenValue-1; j++ {
			r += arr[j] * pow
			pow *= 10
		}
		//fmt.Printf("l = %v, r = %v, arr = %v, i = %v\n", l, r, arr, i)
		if l+r == i {
			fmt.Printf("%v ", i)
			ok = true
		}
	}
	if !ok {
		fmt.Println("INVALID RANGE")
	}
}
