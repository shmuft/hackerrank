package main

import (
	"fmt"
)

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func main() {
	var t int
	var s string
	fmt.Scan(&t)
	var ok = true
	for i := 0; i < t; i++ {
		//fmt.Println(i)
		fmt.Scan(&s)
		revIndex := len(s)
		ok = true
		for j := 1; j < len(s); j++ {
			//one := int(s[j])
			//two := int(s[j-1])
			//fmt.Printf("%v", abs(two - one))
			one := int(s[j])
			two := int(s[j - 1])
			left := abs(one - two)
			one = int(s[revIndex - j])
			two = int(s[revIndex - j - 1])
			right := abs(one - two)
			if left != right {
				fmt.Println("Not Funny")
				ok = false
				break
			}
		}
		if !ok {
			continue
		}
		fmt.Println("Funny")
	}
}
