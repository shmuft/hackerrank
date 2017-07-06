package main

import "fmt"

func main() {
	var n int
	var grade int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&grade)
		if grade >= 38 {
			if ((grade/5+1)*5 - grade) < 3 {
				grade = (grade/5 + 1) * 5
			}
		}
		fmt.Println(grade)
	}
}
