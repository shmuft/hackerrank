package main

import (
	"fmt"
)

func main() {
	var sum, n, k, i, payd, val int

	fmt.Scanf("%v %v", &n, &k)
	sum = 0
	for i=0; i<n; i++{
		fmt.Scanf("%v", &val)
		if i!=k {
			sum += val
		}
	}

	fmt.Scanf("%v", &payd)
	if (sum / 2 == payd) && (sum % 2 == 0){
		fmt.Println("Bon Appetit")
	} else {
		fmt.Println(payd - (sum / 2))
	}
}

