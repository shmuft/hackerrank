package main

import "fmt"

func main() {
	var i, j, n, k, count int
	count = 0
	fmt.Scanf("%v %v", &n, &k)
	var arr = make([]int, n, n)
	for i=0; i<n; i++{
		fmt.Scanf("%v", &arr[i])
	}
	for i=0; i<n-1; i++ {
		for j=i+1; j<n; j++{
            if ((arr[i] + arr[j]) % k == 0){
				count++
			}
		}
	}
	fmt.Println(count)
}



