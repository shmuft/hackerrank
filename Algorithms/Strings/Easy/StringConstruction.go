package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Scanf("%v", &n)
	var s string

	for i:=0; i<n; i++{
		arr := make(map[rune]bool)
		fmt.Scanln(&s)
		for _, j := range s {
			arr[j] = true
		}
		fmt.Println(len(arr))
	}
}

