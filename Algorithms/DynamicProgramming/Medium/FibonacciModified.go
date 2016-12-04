package main

import (
	"fmt"
	"math/big"
)

func main() {
	var t1, t2, n, i int64

	fmt.Scanf("%v %v %v", &t1, &t2, &n)
	tn := big.NewInt(t1)
	tn1 := big.NewInt(t2)
	//tn2 := big.NewInt(0)
	temp := big.NewInt(0)
	for i=2; i<n; i++{
		temp.Mul(tn1, tn1)
		tn.Add(tn,temp)
		tn, tn1 = tn1, tn

	}
	fmt.Println(tn1)
}



