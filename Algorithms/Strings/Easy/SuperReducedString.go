package main

import (
	"fmt"
)

func main() {
	var s string

	fmt.Scanln(&s)

	var flag bool = true
	for flag {
		flag = false
		//fmt.Println(flag)
		for i := 0; i < len(s) -1; i++ {
			if s[i+1] == s[i] {
				flag = true
				switch {
					case i == 0:
						s = s[i+2:]
					case i == len(s)-2:
						//fmt.Println("INSIDDE")
						s = s[:i]
					default:
						//fmt.Printf("i = %v\n", i)
						s = s[:i] + s[i+2:]
				}
				break
			}
		}
	}
	if (len(s) == 0){
		fmt.Println("Empty String")
	} else {
		fmt.Println(s)
	}
}

