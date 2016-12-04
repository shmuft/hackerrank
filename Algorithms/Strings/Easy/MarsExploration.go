package main

import (
	"fmt"
)

func main() {
	var sum int
	var s string
	fmt.Scanf("%v", &s)
	sum = 0
	for i, c := range s{
		switch i%3 {
		case 0:
			if (string(c) != "S") {
				sum++
			}
		case 1:
			if (string(c) != "O") {
				sum++
			}
		case 2:
			if (string(c) != "S") {
				sum++
			}


		}
	}

	fmt.Println(sum)
}

