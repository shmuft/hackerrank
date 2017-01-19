package main

import (
	"fmt"
	"math/big"
)

func main() {
	var n, i int64
	fmt.Scan(&n)
	fact := big.NewInt(1)
	for i = 2; i <= n; i++ {
		fact.Mul(fact, big.NewInt(i))
	}
	fmt.Println(fact)
}
