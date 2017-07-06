package main

import "fmt"

func main() {
	var n, d int
	fmt.Scanf("%d %d", &n, &d)
	array := make([]int, n, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &array[i])
	}
	for i := 0; i < n; i++ {
		fmt.Print(array[(i+d)%n], " ")
	}
}
