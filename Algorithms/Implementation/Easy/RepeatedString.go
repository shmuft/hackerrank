package main

import (
	"fmt"
	"strings"
)

func main() {
	var n, m, count int
	var s string
	fmt.Scanf("%v", &s)
	fmt.Scanf("%v", &n)

	count = strings.Count(s, "a")
	m = n % len(s)

	fmt.Println(count*(n/len(s)) + strings.Count(s[:m], "a"))

}

