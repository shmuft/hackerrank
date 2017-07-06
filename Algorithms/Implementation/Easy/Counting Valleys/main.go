package main

import (
	"fmt"
	"os"
)

func main() {
	var n, level, countValley int
	b := make([]byte, 1, 1)
	fmt.Scanf("%d", &n)
	level = 0
	countValley = 0
	for i := 0; i < n; i++ {
		_, err := os.Stdin.Read(b)
		if err != nil {
			panic("Ups!")
		}
		if string(b) == "U" {
			level++
		}
		if string(b) == "D" {
			if level == 0 {
				countValley++
			}
			level--
		}
	}
	fmt.Println(countValley)
}
