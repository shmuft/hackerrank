package main

import "fmt"

func main() {
	var n, m int
	fmt.Scanf("%d %d", &n, &m)
	array := make([]int64, n, n)
	for i := 0; i < n; i++ {
		array[i] = 0
	}

	var a, b int
	var k int64
	for i := 0; i < m; i++ {
		fmt.Scanf("%d %d %d", &a, &b, &k)
		a--
		b--
		array[a] += k
		if b+1 < n {
			array[b+1] -= k
		}
	}
	var max, sum int64 = 0, 0
	for i := 0; i < n; i++ {
		sum += array[i]
		if sum > max {
			max = sum
		}
	}
	fmt.Println(max)
}
