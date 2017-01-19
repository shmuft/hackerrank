package main

import (
	"fmt"
	"math"
)

func main() {
	var a1, a2, a3, a4, a5, minV, maxV int64
	fmt.Scan(&a1, &a2, &a3, &a4, &a5)
	minV = math.MaxInt64
	maxV = 0

	val := a1+a2+a3+a4
	if val < minV {
		minV = val
	}
	if val > maxV {
		maxV = val
	}

	val = a1+a2+a3+a5
	if val < minV {
		minV = val
	}
	if val > maxV {
		maxV = val
	}

	val = a1+a2+a4+a5
	if val < minV {
		minV = val
	}
	if val > maxV {
		maxV = val
	}

	val = a1+a3+a4+a5
	if val < minV {
		minV = val
	}
	if val > maxV {
		maxV = val
	}

	val = a2+a3+a4+a5
	if val < minV {
		minV = val
	}
	if val > maxV {
		maxV = val
	}

	fmt.Printf("%v %v", minV, maxV)

}
