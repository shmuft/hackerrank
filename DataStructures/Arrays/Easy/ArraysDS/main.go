package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%v", &n)
	array := make([]int, n, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &array[i])
	}

	for i:=n-1; i>=0; i-- {
		fmt.Print(array[i], " ")
	}
}
