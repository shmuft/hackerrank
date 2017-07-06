package main

import (
	"fmt"
)

func main() {
	var n, temp, i, highestScore, lowestScore, min, max int
	fmt.Scanf("%d",&n)
	min = 0
	max = 0
	fmt.Scanf("%d",&temp)
	highestScore = temp
	lowestScore = temp
	for i = 1; i < n; i++ {
		fmt.Scanf("%d",&temp)
//		fmt.Println("max = ", max, ", min=", min, ", hs=", highestScore, ", ls=", lowestScore, ",temp = ", temp)
		if highestScore < temp {
			highestScore = temp
			max++
			//fmt.Println("max = ", max)
		}
		if temp < lowestScore {
			lowestScore = temp
			min++
			//fmt.Println("min = ", min)
		}
	}
	fmt.Println(max, min)
}
