package main

import "fmt"

func main() {
	var x1, x2, v1, v2 int
	fmt.Scanf("%v %v %v %v", &x1, &v1, &x2, &v2)
	if ((v1>v2) && (x1>x2)) || ((v2>v1) && ((x2>x1))){
		fmt.Println("NO")
		return
	}
    if (v1-v2 == 0) {
        fmt.Println("NO")
        return
    } 
	if (x2-x1)%(v1-v2) == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}



