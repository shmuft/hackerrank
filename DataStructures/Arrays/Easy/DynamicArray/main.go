package main

import "fmt"

func main() {
	var n, q int
	fmt.Scanf("%d %d", &n, &q)
	var t, x, y int
	var lastAnswer int = 0
	seqMap := make(map[int][]int)
	var seq int
	for i := 0; i < q; i++ {
		fmt.Scanf("%d %d %d", &t, &x, &y)
		seq = (x ^ lastAnswer) % n
		switch t {
		case 1:
			if _, ok := seqMap[seq]; !ok {
				seqMap[seq] = make([]int, 0, 5)
			}
			seqMap[seq] = append(seqMap[seq], y)
		case 2:
			lastAnswer = seqMap[seq][y%len(seqMap[seq])]
			fmt.Println(lastAnswer)
		}
	}
}
