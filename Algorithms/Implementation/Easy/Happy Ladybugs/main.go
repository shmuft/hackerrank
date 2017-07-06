package main

import "fmt"

func main() {
	var g, n int
	fmt.Scanf("%d", &g)
	for i := 0; i < g; i++ {
		fmt.Scanf("%d", &n)
		letters := make(map[int32]int)

		var temp string
		fmt.Scanf("%s", &temp)
		hasUnder := false
		for _, v := range temp {
			if v == '_' {
				hasUnder = true
			} else {
				letters[v]++
			}
		}
		ok := true
		if hasUnder {
			for _, val := range letters {
				if val == 1 {
					ok = false
					break
				}
			}
		} else {
			var lenSeq int = 1
			letter := temp[0]
			for j := 1; j < len(temp); j++ {
				if temp[j] != letter {
					if lenSeq == 1 {
						ok = false
						break
					} else {
						lenSeq = 1
						letter = temp[j]
						continue
					}
				} else {
					lenSeq++
				}
			}
			if lenSeq == 1 {
				ok = false
			}
		}
		if ok {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
