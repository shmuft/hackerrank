package main

import (
	"fmt"
	"math"
)

func findSmallestNumber(number int64) int64 {
	var i, multip int64
	multip = 1;
	for i=1; i<=number; i++{
		if (math.Mod(float64(multip),float64(i)) != 0){
			for j:=int64(1); j <= i; j++{
				if (math.Mod(float64(multip * j),float64(i))==0){
					multip = multip * j
					break
				}
			}
		}
	}
	return multip
}
func main() {
	var n, i, val int64
	fmt.Scanf("%v", &n)
	for i=0; i<n; i++ {
		fmt.Scanf("%v", &val)
		smallest := findSmallestNumber(val)
		fmt.Println(smallest)
	}
}
