package main

import (
	"fmt"
)

func main() {
	var t, i, j, k int
	var s string
	arr := make([]rune, 100, 100)
	fmt.Scan(&t)
	for i = 0; i < t; i++ {
		fmt.Scan(&s)
		if len(s) <= 1{
			fmt.Println("no answer")
			continue
		}
		for j = 0; j < len(s); j++ {
			arr[j] = rune(s[j])
		}
		for j = len(s) - 1; j > 0; j-- {
			if arr[j - 1] < arr[j] {
				break
			}
		}
		if j == 0 {
			fmt.Println("no answer")
			continue
		}
		k = j
		for j = len(s) - 1; j > 0; j-- {
			if arr[j] > arr[k - 1] {
				break
			}
		}
		arr[k - 1], arr[j] = arr[j], arr[k - 1]

		j = len(s) - 1
		for k < j {
			arr[j], arr[k] = arr[k], arr[j]
			k++
			j--
		}
		lenS := len(s)
		s = ""
		for j := 0; j < lenS; j++ {
			s += string(arr[j])
		}
		fmt.Print(s+"\n")
		//fmt.Print("\n")
	}
}