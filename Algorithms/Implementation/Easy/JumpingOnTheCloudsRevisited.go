package main

import (
	"fmt"
)

func main() {
	var n, k, i, sum, val int

	fmt.Scanf("%v %v", &n, &k)
	var jump int = 0
	sum = 0
	fmt.Scanf("%v", &val)
	if val == 1 {
		sum += 2
	}
	for i = 0; i < n-1; i++{
		fmt.Scanf("%v", &val)
		jump++
		if (jump == k) {
			switch  {
			case val == 1:
				sum += 3
				jump = 0
			case val == 0:
				sum += 1
				jump = 0
			}
		}
		//fmt.Printf("%v %v\n", val, jump)
	}
	sum += 1
	if sum > 100 {
		fmt.Println(0)
	} else {
		fmt.Println(100 - sum)
	}

}

