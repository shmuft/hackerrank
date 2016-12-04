package main

import (
	"fmt"
	"math"
)

func main() {
	var n, i, val, j int

	arr := make(map[int][]int)

	fmt.Scanf("%v", &n)

	for i = 0; i < n ; i++{
		fmt.Scanf("%v", &val)
		arr[val] = append(arr[val], i)
	}
	var identical, dist int

	maxDist := math.MaxInt32
	dist = math.MaxInt32
	for _, val := range arr{
		if len(val) > 1 {
			identical = 0
			for i = 0; i < len(val) - 1; i++ {
				for j = i + 1; j < len(val); j++ {
					switch  {
					case val[j] - val[i] < dist:
						dist = val[j] - val[i]
						identical = 0
					case (val[j] - val[i] == dist):
						identical++
					}
				}
			}
			if (dist < maxDist) && (identical == 0) {
				maxDist = dist
			}
		}
	}
	if maxDist == math.MaxInt32 {
		fmt.Println("-1")
	} else {
		fmt.Println(maxDist)
	}

}

