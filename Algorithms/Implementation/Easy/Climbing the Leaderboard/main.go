package main

import (
	"fmt"
)

func main() {
	var last, temp, n, m, aliceScore, i, j int
	fmt.Scanf("%d", &n)
	scores := make([]int, 0, 0)
	last = -1
	for i = 0; i < n; i++ {
		fmt.Scanf("%d", &temp)
		if temp != last {
			scores = append(scores, temp)
			last = temp
		}
	}

	//fmt.Println(scores)
	fmt.Scanf("%d", &m)
	j = len(scores) - 1;
	for i = 0; i < m; i++ {
		fmt.Scanf("%d", &aliceScore)
		if j >= 0 {
			for ; j >= 0; j-- {
				//fmt.Println("j = ",j,"scores[j]=",scores[j], "aliceScore=",aliceScore)
				if aliceScore == scores[j] {
					fmt.Println(j + 1)
					break
				}
				if aliceScore < scores[j] {
					fmt.Println(j + 2)
					break
				}
				if aliceScore > scores[0] {
					fmt.Println(1)
					j--;
					break
				}

			}
		} else {
			fmt.Println(1)
		}
	}
}
