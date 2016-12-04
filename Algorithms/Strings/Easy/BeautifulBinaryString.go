package main

import "fmt"

func main() {
	var n int
	var b string

	fmt.Scanf("%v", &n)
	arr := make([]int, n, n)
	fmt.Scanln(&b)

	for i := 0; i < len(b); i++ {
		if b[i] == 48 {
			arr[i] = 0
		} else {
			arr[i] = 1
		}
	}
	var sum int = 0
	// fmt.Println(arr)
	for i := 0; i < len(arr)-2; i++ {
		if (arr[i] == 0) && (arr[i+1] == 1) && (arr[i+2] == 0) {
			arr[i+2] = 1
			sum++
		}
	}

	fmt.Println(sum)

}

