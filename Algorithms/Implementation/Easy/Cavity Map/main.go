package main

import "fmt"

//https://www.hackerrank.com/challenges/cavity-map
func main() {
	var n int
	fmt.Scanf("%d", &n)
	squareMap := make([][]byte, n, n)
	for i := 0; i < n; i++ {
		squareMap[i] = make([]byte, n, n)
		fmt.Scanf("%s", &squareMap[i])
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i > 0 && j > 0 && i < n-1 && j < n-1 {
				if squareMap[i][j] > squareMap[i][j-1] &&
					squareMap[i][j] > squareMap[i][j+1] &&
					squareMap[i][j] > squareMap[i-1][j] &&
					squareMap[i][j] > squareMap[i+1][j] {

					fmt.Print("X")
				} else {
					fmt.Print(squareMap[i][j] - 48)
				}
			} else {
				fmt.Print(squareMap[i][j] - 48)
			}
		}
		fmt.Println()
	}

}
