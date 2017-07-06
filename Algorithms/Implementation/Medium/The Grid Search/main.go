package main

import "fmt"

//https://www.hackerrank.com/challenges/the-grid-search
func isFoundPatternInGrid(indexRow, indexColumn int, grid, pattern [][]byte) bool {
	for i := 0; i < len(pattern); i++ {
		for j := 0; j < len(pattern[i]); j++ {
			if grid[indexRow+i][indexColumn+j] != pattern[i][j] {
				return false
			}
		}
	}
	return true
}

func FindPattern(r, c, r1, c1 int, grid, pattern [][]byte) bool {
	for indexRow := 0; indexRow <= r-r1; indexRow++ {
		for indexColumn := 0; indexColumn <= c-c1; indexColumn++ {
			if grid[indexRow][indexColumn] == pattern[0][0] {
				if ok := isFoundPatternInGrid(indexRow, indexColumn, grid, pattern); ok {
					return true
				}
			}
		}
	}
	return false
}
func main() {
	var t, r, c, r1, c1 int
	fmt.Scanf("%d", &t)
	for i := 0; i < t; i++ {

		//Input data
		fmt.Scanf("%d", &r)
		fmt.Scanf("%d", &c)
		grid := make([][]byte, r, r)
		for j := 0; j < r; j++ {
			grid[j] = make([]byte, c, c)
			fmt.Scanf("%s", &grid[j])
		}
		fmt.Scanf("%d", &r1)
		fmt.Scanf("%d", &c1)
		pattern := make([][]byte, r1, r1)
		for j := 0; j < r1; j++ {
			pattern[j] = make([]byte, c1, c1)
			fmt.Scanf("%s", &pattern[j])
		}

		//Find pattern in grid
		ok := FindPattern(r, c, r1, c1, grid, pattern)
		if ok {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
