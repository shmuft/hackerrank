package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	arrayStrings := make([]string, n, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%s", &arrayStrings[i])
	}
	var q int
	fmt.Scanf("%d", &q)
	var s string
	for i := 0; i < q; i++ {
		fmt.Scanf("%s", &s)
		var sum int = 0
		for j := 0; j < n; j++ {
			if len(arrayStrings[j]) == len(s) {
				var ok bool = true
				for k := 0; k < len(s); k++ {
					if arrayStrings[j][k] != s[k] {
						ok = false
						break
					}
				}
				if ok {
					sum++
				}
			}
		}
		fmt.Println(sum)
	}

}
