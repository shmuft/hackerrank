package main

import "fmt"

func main() {
	var t, dollars, cost, costWrapper int
	fmt.Scanf("%d", &t)
	for i := 0; i < t; i++ {
		fmt.Scanf("%d", &dollars)
		fmt.Scanf("%d", &cost)
		fmt.Scanf("%d", &costWrapper)
		amountChocolates := dollars / cost
		chocolateWrappers := amountChocolates
		for chocolateWrappers >= costWrapper {
			amountChocolates += chocolateWrappers / costWrapper
			chocolateWrappers = chocolateWrappers/costWrapper + chocolateWrappers%costWrapper
		}
		fmt.Println(amountChocolates)
	}
}
