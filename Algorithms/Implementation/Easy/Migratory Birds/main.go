package main

import "fmt"

func main() {
	var n, temp, max, keyMax int
	fmt.Scanf("%d", &n)
	arr := make(map[int]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &temp)
		_, ok := arr[temp]
		if !ok {
			arr[temp] = 1
		} else {
			arr[temp]++
		}
	}
	max = 0;
	keyMax = -1;
	for key, val := range arr {
		if val > max {
			max = val
			keyMax = key
		} else if val == max && key < keyMax {
			max = val
			keyMax = key
		}
	}
	fmt.Println(keyMax)
}
