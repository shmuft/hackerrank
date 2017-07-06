package main

import (
	"fmt"
)

func CoordinatesOfElementInSequence(m, n, positionInSequence int) (y, x int) {
	switch {
	case positionInSequence < m:
		return positionInSequence, 0
	case positionInSequence >= m && positionInSequence < m+n-2:
		return m - 1, positionInSequence - m + 1
	case positionInSequence >= m+n-2 && positionInSequence < 2*m+(n-2):
		return (m - 1) - (positionInSequence - (m + n - 2)), n - 1
	default:
		return 0, 2*m + 2*n - 4 - positionInSequence
	}
}

func main() {
	var m, n, r int
	fmt.Scanf("%d", &m)
	fmt.Scanf("%d", &n)
	fmt.Scanf("%d", &r)
	matrix := make([][]int, m, m)
	matrix2 := make([][]int, m, m)
	for i := 0; i < m; i++ {
		matrix[i] = make([]int, n, n)
		matrix2[i] = make([]int, n, n)
		for j := 0; j < n; j++ {
			fmt.Scanf("%d", &matrix[i][j])
		}
	}
	var lenDiag int
	if n < m {
		lenDiag = n / 2
	} else {
		lenDiag = m / 2
	}
	m1, n1 := m, n
	var y1, x1, y, x, lenSequence int
	for i := 0; i < lenDiag; i, m1, n1 = i+1, m1-2, n1-2 {
		lenSequence = m1*2 + n1*2 - 4
		r1 := r % lenSequence
		for j := 0; j < lenSequence; j++ {
			y1, x1 = CoordinatesOfElementInSequence(m1, n1, (j+r1)%lenSequence)
			y, x = CoordinatesOfElementInSequence(m1, n1, j)
			matrix2[y1+i][x1+i] = matrix[y+i][x+i]
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(matrix2[i][j], " ")
		}
		fmt.Println()
	}
}
