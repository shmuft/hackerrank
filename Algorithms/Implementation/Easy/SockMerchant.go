package main

import "fmt"

func main() {
	var n, i, val int
	fmt.Scanf("%v", &n)
	arr := make(map[int]int)
	for i = 0; i < n; i++ {
		fmt.Scanf("%v", &val)
		_, ok := arr[val]

		if ok {
			arr[val]++
		} else {
			arr[val] = 1
		}
	}
	var sum int
	sum = 0
	for _, val := range arr {
		sum += val >> 1
	}
	fmt.Println(sum)
}

