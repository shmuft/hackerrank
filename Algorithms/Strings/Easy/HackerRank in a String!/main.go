package main

import "fmt"

func main() {
	var q int
	fmt.Scanf("%v", &q)
	hack := "hackerrank"
	var str string
	for i := 0; i < q; i++ {
		cur := 0
		fmt.Scanln(&str)
		for _, val := range str {
			if byte(val) == hack[cur] {
				cur++
				if cur == 10 {
					break
				}
			}
		}
		if cur == 10 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
