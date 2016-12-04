package main

import (
	"fmt"
)

const sizeN = 20
var grid [sizeN][sizeN]int
var ch = make(chan int, 4)

func upDown(){
	prod := 1
	var maxProd = 0
	for i := 0; i<sizeN; i++{
		for j:=0; j<sizeN-3; j++{
			prod = 1
			for k:=0; k<4; k++{
				prod *= grid[j+k][i]
			}
			if prod>maxProd {
				maxProd = prod
			}
		}
	}
	ch <- maxProd
}

func leftRight(){
	prod := 1
	var maxProd = 0
	for i := 0; i<sizeN; i++{
		for j:=0; j<sizeN-3; j++{
			prod = 1
			for k:=0; k<4; k++{
				prod *= grid[i][j+k]
			}
			if prod>maxProd {
				maxProd = prod
			}
		}
	}
	ch <- maxProd
}

func mainDiag(){
	prod := 1

	var maxProd = 0
	for i := 0; i<=sizeN-4; i++{
		for j:=0; j<=sizeN-4; j++{
			prod = 1
			for k:=0; k<4; k++{
				prod *= grid[i+k][j+k]
			}
			if prod>maxProd {
				maxProd = prod
			}
		}
	}
	ch <- maxProd
}

func notMainDiag(){
	prod := 1

	var maxProd = 0
	for i := 0; i<=sizeN-4; i++{
		for j:=3; j<sizeN; j++{
			prod = 1
			for k:=0; k<4; k++{
				prod *= grid[i+k][j-k]
			}
			if prod>maxProd {
				maxProd = prod
			}
		}
	}
	ch <- maxProd
}

func main() {
	for i:=0; i<sizeN; i++{
		for j:=0; j<sizeN; j++{
			fmt.Scanf("%v", &grid[i][j])
		}
	}

	i := 0
	go upDown()
	go leftRight()
	go mainDiag()
	go notMainDiag()
	i = 0
	var maxProd int = 0
	for prod := range ch{
		if prod > maxProd {
			maxProd = prod
		}
		i++
		if i == 4 {
			close(ch)
		}
	}
	fmt.Println(maxProd)
}
