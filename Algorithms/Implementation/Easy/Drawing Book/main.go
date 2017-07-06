package main

import "fmt"

func main() {
	var n, p int
	fmt.Scanf("%d", &n)
	fmt.Scanf("%d", &p)
	if p/2 < (n-p)/2 {
		fmt.Println(p / 2)
	} else {
		fmt.Println((n - p) / 2)
	}
}
