package main

import (
	"fmt"
	"os"
	"sort"
)

type flavor struct {
	cost int
	id   int
}

type flavors []flavor

func (f flavors) Len() int           { return len(f) }
func (f flavors) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }
func (f flavors) Less(i, j int) bool { return f[i].cost < f[j].cost }

func Round(a, div int) int {
	if a%div != 0 {
		if float64(a%div)/float64(div) > 0.4 {
			return a/div + 1
		}
	}
	return a / div
}

func BSearch(array flavors, elem int, cantUse int) (int, bool) {
	var left, right int = 0, len(array)-1

	for itemIndex := left + (right-left)/2; ; itemIndex = (right-left)/2 + left {
		//fmt.Println(left, right)
		switch {
		case elem == array[itemIndex].cost && itemIndex != cantUse:
			return itemIndex, true
		case left >= right:
			return 0, false
		case elem < array[itemIndex].cost:
			right = itemIndex
		case elem >= array[itemIndex].cost:
			left = itemIndex + 1
		}
	}
}

func main() {
	var t int
	fmt.Fscanf(os.Stdin, "%d", &t)
	var m, n int
	for i := 0; i < t; i++ {
		fmt.Scanf("%d", &m)
		fmt.Scanf("%d", &n)
		array := make(flavors, n, n)
		var temp int
		for j := 0; j < n; j++ {
			fmt.Fscanf(os.Stdin, "%d", &temp)
			array[j].cost = temp
			array[j].id = j
		}
		sort.Sort(array)
		for j := 0; j < n; j++ {
			secondFlavour := m - array[j].cost
			if indexItem, ok := BSearch(array, secondFlavour, j); ok {
				if array[j].id < array[indexItem].id {
					fmt.Printf("%d %d\n", array[j].id+1, array[indexItem].id+1)
				} else {
					fmt.Printf("%d %d\n", array[indexItem].id+1, array[j].id+1)
				}
				break
			}
		}
	}
}
