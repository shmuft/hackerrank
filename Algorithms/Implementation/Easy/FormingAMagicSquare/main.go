package main

import (
	"fmt"
)

func isSquareMagic(square [][]int) (bool, []int) {
	arraySum := make([]int, 8, 8)
	for i := 0; i < 3; i++ {
		arraySum[0] += square[0][i]
		arraySum[1] += square[1][i]
		arraySum[2] += square[2][i]
		arraySum[3] += square[i][0]
		arraySum[4] += square[i][1]
		arraySum[5] += square[i][2]
		arraySum[6] += square[i][i]
		arraySum[7] += square[i][2-i]
	}
	var sum int = arraySum[0]
	for i := 1; i < 8; i++ {
		if arraySum[i] != sum {
			return false, arraySum
		}
	}
	return true, arraySum
}

func mod(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func mirrorMatrix(matrix [][]int) {
	for i := 0; i < 3; i++ {
		matrix[i][0], matrix[i][2] = matrix[i][2], matrix[i][0]
	}
}

func rotateMatrix(matrix [][]int) [][]int {
	res := make([][]int, 3, 3)
	res[0] = make([]int, 3, 3)
	res[1] = make([]int, 3, 3)
	res[2] = make([]int, 3, 3)

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			res[i][j] = matrix[j][i]
		}
	}

	for i := 0; i < 3; i++ {
		res[0][i], res[2][i] = res[2][i], res[0][i]
	}
	return res
}

func calculate(workingMatrix, comparableMatrix [][]int) int {
	var cost int = 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			cost += mod(workingMatrix[i][j] - comparableMatrix[i][j])
		}
	}
	return cost
}

func main() {
	magicSquare := make([][]int, 3, 3)
	for i := 0; i < 3; i++ {
		magicSquare[i] = make([]int, 3, 3)
		for j := 0; j < 3; j++ {
			fmt.Scanf("%v", &magicSquare[i][j])
		}
	}

	magic := make([][]int, 3, 3)
	magic[0] = make([]int, 3, 3)
	magic[1] = make([]int, 3, 3)
	magic[2] = make([]int, 3, 3)

	magic[0][0] = 8
	magic[0][1] = 1
	magic[0][2] = 6
	magic[1][0] = 3
	magic[1][1] = 5
	magic[1][2] = 7
	magic[2][0] = 4
	magic[2][1] = 9
	magic[2][2] = 2

	var minCost int = 100

	for i := 0; i < 4; i++ {
		if cost := calculate(magic, magicSquare); cost < minCost {
			minCost = cost
		}
		magicSquare = rotateMatrix(magicSquare)
	}

	mirrorMatrix(magicSquare)
	for i := 0; i < 4; i++ {
		if cost := calculate(magic, magicSquare); cost < minCost {
			minCost = cost
		}
		magicSquare = rotateMatrix(magicSquare)
	}

	fmt.Println(minCost)
}
