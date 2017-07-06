package main

import "fmt"

func main() {
	var n, d, m int
	fmt.Scanf("%d", &n)
	arr := make([]int, n, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &arr[i])
	}
	fmt.Scanf("%d %d", &d, &m)
	var sum, count int
	sum = 0
	count = 0
	for i := 0; i < m; i++ {
		sum += arr[i]
	}
	if sum == d {
		count++
	}
	for i := m; i < n; i++ {
		sum -= arr[i-m]
		sum += arr[i]
		if sum == d {
			count++
		}
	}
	fmt.Println(count)
}
