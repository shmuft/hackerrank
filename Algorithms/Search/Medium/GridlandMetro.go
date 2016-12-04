package main

import (
	"fmt"
)

type arrType map[intType][]segment
type intType int64

var arr arrType
var n, m, k, r, c1, c2, i, sum, sum2 intType
//var ok bool
//var sL, sR,,
var sI, STemp segment

type segment struct {
	c1, c2 intType
}

func isSubset(a, b segment) bool {
	return ((a.c1 <= b.c1) && (a.c2 >= b.c2)) || ((b.c1 <= a.c1) && (b.c2 >= a.c2))
}

func isOverlap(a, b segment) bool {
	return (b.c1 < a.c1 && a.c1 <= b.c2 && b.c2 < a.c2) || (a.c1 < b.c1 && b.c1 <= a.c2 && a.c2 < b.c2)
}

func union(a, b segment) segment {
	if a.c1 < b.c1 {
		STemp.c1 = a.c1
	} else {
		STemp.c1 = b.c1
	}

	if a.c2 > b.c2 {
		STemp.c2 = a.c2
	} else {
		STemp.c2 = b.c2
	}

	return STemp
}

//func mySum(arr arrType, n, m intType, maxSize int) intType {
//	for keyX, x := range arr {
//		ok = true
//		startArr := x
//		endArr := make([]segment, 0, len(x))
//		for ok {
//			for _, y := range startArr {
//				if len(endArr) == 0 {
//					endArr = append(endArr, y)
//				} else {
//					for key, z := range endArr {
//						sI = segment{c1: -1, c2: -1}
//						if y.c1 > z.c1 {
//							sL = z
//							sR = y
//						} else {
//							sL = y
//							sR = z
//						}
//						if sR.c1 <= sL.c2 {
//							sI.c1 = sL.c1
//							if sR.c2 <= sL.c2 {
//								sI.c2 = sL.c2
//							} else {
//								sI.c2 = sR.c1
//							}
//						}
//
//						if sI.c1 == -1 && sI.c2 == -1 {
//							endArr = append(endArr, y)
//						} else {
//							endArr[key] = sI
//						}
//					}
//				}
//			}
//			if len(endArr) != len(startArr) {
//				ok = true
//				startArr = endArr
//				endArr = nil
//			} else {
//				arr[keyX] = endArr
//				ok = false
//			}
//		}
//	}
//	sum = 0
//	sum += n * m
//	for _, s := range arr {
//		sum2 = 0
//		for _, x := range s {
//			sum2 += x.c2 - x.c1 + 1
//		}
//		//  fmt.Println(sum2)
//		sum -= sum2
//	}
//	return sum
//}

func main() {
	fmt.Scanf("%v %v %v", &n, &m, &k)
	var overlap bool
	arr = make(map[intType][]segment)
	for i = 0; i < k; i++ {
		fmt.Scanf("%v %v %v", &r, &c1, &c2)
		if c1 > c2 {
			c1, c2 = c2, c1
		}
		sI = segment{c1:c1, c2:c2}
		overlap = false
		for key, x := range arr[r] {
			if (isSubset(sI, x) || isOverlap(sI, x)) {
				overlap = true
				arr[r][key] = union(sI, x)
				break
			}
		}
		if !overlap {
			arr[r] = append(arr[r], sI)
		}
	}
	//var maxSize int = 0
	//for i = 0; i < k; i++ {
	//	if len(arr[i]) > maxSize {
	//		maxSize = len(arr[i])
	//	}
	//}
	sum = 0
	sum += n * m
	for _, s := range arr {
		sum2 = 0
		for _, x := range s {
			sum2 += x.c2 - x.c1 + 1
		}
		//  fmt.Println(sum2)
		sum -= sum2
	}
	fmt.Println(sum)
	//fmt.Println(mySum(arr, n, m, maxSize))

}

