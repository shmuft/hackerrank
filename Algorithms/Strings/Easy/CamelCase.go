package main

import (
	"fmt"
	"strings"
)

func main() {
	var sum int
	var s string
	fmt.Scanf("%v", &s)

	for _, i := range s{
		if string(i) != strings.ToLower(string(i)) {
			sum++
		}
	}
	sum++

	fmt.Println(sum)
}

