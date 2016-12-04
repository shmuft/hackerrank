package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanf("%v", &n)

	arr:= make([]int, n, n)

	for i:=0; i< n; i++{
		fmt.Scanf("%v", &arr[i])
	}
	if n == 2{
		if arr[0]%2 == 1 || arr[1] % 2 == 1 {
			fmt.Println("NO")
		} else {
			fmt.Println(0)
		}
	}else {
		var sum int = 0
		for i := 0; i < n - 1; i++ {
			if arr[i] % 2 == 1 {
				arr[i]++
				arr[i + 1]++
				sum++
			}
		}
		var flag bool = true
		for i:=0; i<n-1; i++{
			if arr[i]%2 == 1 || arr[i+1] % 2 == 1 {
				fmt.Println("NO")
				flag = false
				break
			}
		}
		if flag {
			fmt.Println(2 * sum)
		}
	}
}

