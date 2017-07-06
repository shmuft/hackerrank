package main

import (
	"fmt"
)

func main() {
	var n, k int
	var rq, cq int

	fmt.Scanf("%d", &n)
	fmt.Scanf("%d", &k)
	fmt.Scanf("%d", &rq)
	fmt.Scanf("%d", &cq)

	var minLeft, minRight, minUp, minDowm int
	var minLeftMain, minRightMain int
	var minLeftNotMain, minRightNotMain int

	minLeft = cq - 1
	minRight = n - cq
	minDowm = rq - 1
	minUp = n - rq

	//y = kx + b, where k = 1; rq = cq + b; b = rq - cq; y = x + rq - cq

	var b1 int = rq - cq
	if b1 > 0 {
		minLeftMain = cq - 1
		minRightMain = n - b1 - minLeftMain - 1
	} else {
		minLeftMain = rq - 1
		minRightMain = n + b1 - minLeftMain - 1
	}

	//y = kx + b, where k = -1; rq = -cq + b; b = rq + cq; y = -x + rq + cq
	//y = -x + b; y = -1 + b
	//1 = -x + b; x = b - 1

	//n = -x + b; x = b - n
	//y = -n + b
	var b2 int = rq + cq
	if -1+b2 < n {
		minLeftNotMain = (-1 + b2) - rq
		minRightNotMain = b2 - 1 - cq
	} else {
		minLeftNotMain = cq - (b2 - n)
		minRightNotMain = rq - (-n + b2)
	}

	//fmt.Println("------------------------------INITIAL")
	//fmt.Println(minLeft, minRight)
	//fmt.Println(minUp, minDowm)
	//fmt.Println(minLeftMain, minRightMain)
	//fmt.Println(minLeftNotMain, minRightNotMain)

	var r, c int
	for i := 0; i < k; i++ {
		fmt.Scanf("%d %d", &r, &c)
		switch {
		case r == c+b1:
			if r < rq && rq-r-1 < minLeftMain {
				minLeftMain = rq - r - 1
			} else if r > rq && r-rq-1 < minRightMain {
				minRightMain = r - rq - 1
			}
		case r == -c+b2:
			if r < rq && rq-r-1 < minRightNotMain {
				minRightNotMain = rq - r - 1
			} else if r > rq && r-rq-1 < minLeftNotMain {
				minLeftNotMain = r - rq - 1
			}
		case r == rq:
			if c < cq && cq-c-1 < minLeft {
				minLeft = cq - c - 1
			} else if cq < c && c-cq-1 < minRight {
				minRight = c - cq - 1
			}
		case c == cq:
			if r < rq && rq-r-1 < minDowm {
				minDowm = rq - r - 1
			} else if rq < r && r-rq-1 < minUp {
				minUp = r - rq - 1
			}
		}
		//fmt.Println("------------------------------")
		//fmt.Println(minLeft, minRight)
		//fmt.Println(minUp, minDowm)
		//fmt.Println(minLeftMain, minRightMain)
		//fmt.Println(minLeftNotMain, minRightNotMain)
	}

	fmt.Println(minLeft + minRight + minUp + minDowm + minLeftMain + minRightMain + minLeftNotMain + minRightNotMain)
}
