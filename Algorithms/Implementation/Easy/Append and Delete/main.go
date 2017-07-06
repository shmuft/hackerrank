package main

import (
	"fmt"
)

func main() {
	var s, t string
	var k int
	fmt.Scan(&s)
	fmt.Scan(&t)
	fmt.Scan(&k)
	i := 0
	for i = 0; i < len(s) && i < len(t); i++ {
		if (s[i] != t[i]) {
			break
		}
	}
	ok := false

	if len(t) < len(s) {
		num := k - (len(s) - i)
		//fmt.Println(num)
		if num%2 == 0 && num > 0 || (k - num - 1 == len(t)){
			ok = true
		}
	}
	if len(s) < len(t) {
		num := k - (len(t) - i)
		//fmt.Println(num)
		if num%2 == 0 && num > 0 {
			ok = true
		}
	}

	if len(s) == len(t) {
		switch {
		case i == len(s):
			if k == len(s)*2+1 {
				ok = true
			}
			if (k < len(s)*2 && k%2 == 0) {
				ok = true
			}
		case i == 0:
			//fmt.Println(i)
			//fmt.Println(len(s))
			if k*2 <= len(s) && k%2 == 0 {
				ok = true
			}
			if k >= len(s)*2 {
				ok = true
			}
		default:
			var num int

			num = (len(s) - i) * 2
			if num <= k && num%2 == 0 {
				ok = true
			}
		}
	}

	if ok {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
