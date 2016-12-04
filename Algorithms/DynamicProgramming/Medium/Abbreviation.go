package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	fmt.Scanf("%v", &n)
	var a, b string
	for i := 0; i < n; i++ {
		fmt.Scanln(&a)
		fmt.Scanln(&b)
		//fmt.Printf("i = %v\n", i)
		if len(a) == len(b) {
			if strings.ToUpper(a) == b {
				fmt.Println("YES")
				continue
			}
		}
		substr := a
		ok := true
		for j := 0; j < len(b); j++ {
			ind := strings.IndexAny(substr, strings.ToLower(string(b[j]))+string(b[j]))
			if ind != -1 {
				temp := substr[:ind]
				if strings.ToLower(temp) == temp {
					substr = substr[ind+1:]
				} else {
					ok2 := false
					k := 0
					if j != 0 {
						for ; k < len(temp); k++ {
							if string(substr[k]) == string(b[j-1]) {
								// fmt.Printf("substr[%v] = %v\n", k, string(substr[k]))
								// fmt.Println(substr)
								// fmt.Println()
								// fmt.Printf("k! = %v\n", k)
								ok2 = true
							} else {
								// fmt.Printf("k!! = %v\n", k)
								break
							}
						}
					}
					// fmt.Printf("k = %v\n", k)
					if ok2 {
						k--
						substr = substr[ind+1+k:]
						// fmt.Sprintln(substr)
					} else {
						// fmt.Println("\nERROR!!!")
						// fmt.Println(substr)
						// fmt.Println(string(b[j]))
						// fmt.Println(string(b[j-1]))
						// fmt.Println(temp)
						// fmt.Println(ind)
						// fmt.Println(123)

						ok = false
						break
					}
				}
			} else {
				// fmt.Println(substr)
				// fmt.Println(456)
				ok = false
				break
			}
		}
		if len(substr) != 0 && ok {
			if strings.ToLower(substr) == substr {
				ok = true
			} else {
				ok2 := false
				k := 0
				for ; k < len(substr)-1; k++ {
					if string(substr[k]) == strings.ToLower(string(b[len(b)-1])) {
						ok2 = true
					} else {
						break
					}
				}
				if string(substr[len(substr)-1]) == string(b[len(b)-1]) {
					ok2 = true
				} else {
					ok2 = false
				}

				if ok2 {
					ok = true
				} else {
					ok = false
				}
			}
		}
		if ok {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}

	}
}

