package main

import "fmt"

func main() {
	var n, temp int
	fmt.Scanf("%d", &n)
	arr := make(map[int]int)
	for i := 1; i <= n; i++ {
		fmt.Scanf("%d", &temp)
		arr[temp] = i
	}
	for i := 1; i <= n; i++ {
		fmt.Println(arr[arr[i]])
	}

}
