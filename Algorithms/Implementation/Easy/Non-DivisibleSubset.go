package main

import "fmt"

func main() {
	var n, val, k, i int64
	fmt.Scanf("%v %v", &n, &k)
	arr := make([]int64, k, k)
	for i = 0; i < k; i++ {
		arr[i] = 0
	}

	for i = 0; i < n; i++ {
		fmt.Scanf("%v", &val)
		arr[val%k]++
	}

	if k == 0 || k == 1 {
		fmt.Println(1)
		return
	}
	// fmt.Println(arr)
	var sum int64 = 0
	for i = 1; i < k-i; i++ {
		if arr[i] > arr[k-i] {
			sum += arr[i]
		} else {
			sum += arr[k-i]
		}
	}

	if k%2 == 0 && len(arr) > 2 {
		sum++
	}
	if arr[0] != 0 {
		sum++
	}
	if len(arr) == 2 {
		if arr[1] > 0 {
			sum++
		}
	}
	fmt.Println(sum)

}

