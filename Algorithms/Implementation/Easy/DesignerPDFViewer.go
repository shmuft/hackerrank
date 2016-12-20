package main

import "fmt"

func main() {
	const n = 26
	arr := make([]int, n, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	var s string
	fmt.Scanln(&s)
	max := 0
	for _, v := range s {
		if arr[int(v) - 97] > max {
			max = arr[int(v) - 97]
		}
	}
	fmt.Println(max * len(s))
}

