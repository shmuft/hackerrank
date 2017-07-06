package main

import "fmt"

func main() {
	var n, height, maxHeight int
	fmt.Scan(&n)
	maxHeight = -1
	candleHeight := make(map[int]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&height)
		_, ok := candleHeight[height]
		if !ok {
			candleHeight[height] = 1
		} else {
			candleHeight[height]++
		}
		if candleHeight[height] > maxHeight {
			maxHeight = candleHeight[height]

		}
	}
	fmt.Println(maxHeight)

}
