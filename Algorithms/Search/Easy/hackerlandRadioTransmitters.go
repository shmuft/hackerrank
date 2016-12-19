package main

import (
	"fmt"
	"sort"
)

func findNext(left int, k int, arr []int) int {
	i := left + 1
	for ; i < len(arr) && arr[i] - arr[left] <= k; i++ {
		//fmt.Printf("arr[%v] - arr[%v] = %v; i = %v\n", arr[i], arr[left], arr[i]- arr[left], i)
	}
	return i-1
}

func main() {
	var n, k, temp int
	arr := make([]int, n, n)
	arrTemp := make(map[int]bool)
	fmt.Scan(&n, &k)
	for i := 0; i < n; i++ {
		fmt.Scan(&temp)
		_, ok := arrTemp[temp]
		if !ok {
			arr = append(arr, temp)
			arrTemp[temp] = true
		}
		//fmt.Println(i)
	}

	sort.Ints(arr)
	//fmt.Println(arr)
	sum := 0
	if len(arr) == 1 {
		sum = 1
	} else {
		var left, right int
		left = -1
		var ok bool = true
		for ok {
			right = findNext(left + 1, k, arr)
			//fmt.Printf("right = %v\n", right)
			left = findNext(right, k, arr)
			//fmt.Printf("left = %v\n", left)
			sum++
			if left == len(arr) -1 {
				ok = false
			}
			//fmt.Println("there")
		}
	}
	fmt.Println(sum)
}

