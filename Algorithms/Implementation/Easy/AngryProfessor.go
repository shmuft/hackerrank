package main

import "fmt"

func main() {
	var i, t, n, k, j, val, good int
	fmt.Scanf("%v", &t)
	for i=0; i<t; i++{
		good = 0
		fmt.Scanf("%v %v", &n, &k)
		for j=0; j<n; j++{
			fmt.Scanf("%v", &val)
			if (val <= 0 ){
				good++
			}
		}
		if (good >= k){
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
		}
	}
}



