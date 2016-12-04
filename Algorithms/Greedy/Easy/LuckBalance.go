package main

import (
	"fmt"
	"math"
)

func main() {
	var n, k, l, t int

	fmt.Scanf("%v %v", &n, &k)
	arrImp := make([]int, 0, n)
	arrNotImp := make([]int, 0, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%v %v", &l, &t)
		if t == 1 {
			arrImp = append(arrImp, l)
		} else {
			arrNotImp = append(arrNotImp, l)
		}
	}
	// fmt.Println(arrImp)
	// fmt.Println(arrNotImp)
	// fmt.Println()
	var foundK int = 0
	var x, sum int
	sum = 0
	var lowest int

	for (len(arrImp) > k) && (len(arrImp) != 0) {
		lowest = math.MaxInt32
		for i := 0; i < len(arrImp); i++ {
			if arrImp[i] < lowest {
				lowest = arrImp[i]
				x = i
			}
		}
		sum -= lowest
		if x+1 < len(arrImp) {
			arrImp = append(arrImp[:x], arrImp[x+1:]...)
		} else {
			arrImp = arrImp[:x]
		}
		// fmt.Println(arrImp)
		foundK++
	}
	// fmt.Println(sum)

	// lowest = math.MaxInt32
	// for (len(arrNotImp)-foundK > k) && (foundK < len(arrNotImp)) {
	// 	for i := 0; i < len(arrNotImp); i++ {
	// 		if arrNotImp[i] < lowest {
	// 			lowest = arrNotImp[i]
	// 			x = i
	// 		}
	// 	}
	// 	sum -= lowest
	// 	if x+1 < len(arrNotImp) {
	// 		arrImp = append(arrNotImp[:x], arrNotImp[x+1:]...)
	// 	} else {
	// 		arrImp = arrNotImp[:x]
	// 	}
	// 	// fmt.Println(arrNotImp)
	// 	foundK++
	// }

	for i := 0; i < len(arrImp); i++ {
		sum += arrImp[i]
	}

	for i := 0; i < len(arrNotImp); i++ {
		sum += arrNotImp[i]
	}
	fmt.Println(sum)

}

