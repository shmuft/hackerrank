package main

import (
	"fmt"
	"math"
)

func main() {
	var edges map[int]map[int]int
	var n, m, i, x, y, s, r int
	edges = make(map[int]map[int]int)
	subGraph := make(map[int]map[int]int)
	nodes := make(map[int]bool)
	fmt.Scanf("%v %v ", &n, &m)

	for i = 0; i < m; i++ {
		fmt.Scanf("%v %v %v ", &x, &y, &s)

		_, ok := edges[x]
		if ok {
			edges[x][y] = s
		} else {
			edges[x] = make(map[int]int)
			edges[x][y] = s
		}

		_, ok = edges[y]
		if ok {
			edges[y][x] = s
		} else {
			edges[y] = make(map[int]int)
			edges[y][x] = s
		}
	}
	fmt.Scanf("%v", &r)
	last := r
	nodes[last] = true
	var weight int
	weight = 0
	var lowestWeight int
	lowestWeight = math.MaxInt32
	for i := 0; i < n-1; i++ {
		//собираем рёбра, которые нужно посмотрет
		//ищём самый малый вес
		lowestWeight = math.MaxInt32
		for key, val := range edges[last] {
			_, ok2 := nodes[key]
			if !ok2 {
				_, ok := subGraph[last]
				if ok {
					subGraph[last][key] = val
				} else {
					subGraph[last] = make(map[int]int)
					subGraph[last][key] = val
				}
			}
		}
		x = 0
		y = 0
		for key, val := range subGraph {
			for key2, val2 := range val {
				_, ok := nodes[key2]
				if !ok {
					switch {
					case val2 < lowestWeight:
						x = key
						y = key2
						lowestWeight = val2
						// case val2 == lowestWeight:
						// 	if (key + val2 + key2) < (x + lowestWeight + y) {
						// 		x = key
						// 		y = key2
						// 		lowestWeight = val2
						// 	}
					}
				}
			}
		}
		weight += lowestWeight
		nodes[y] = true
		last = y

		//убираем из ребёр "посмотреть" найденное ребро и из основного дерева ребро
		delete(subGraph[x], y)
		if len(subGraph[x]) == 0 {
			delete(subGraph, x)
		}
		delete(edges[x], y)
		if len(edges[x]) == 0 {
			delete(edges, x)
		}
		delete(edges[y], x)
		if len(edges[y]) == 0 {
			delete(edges, y)
		}
		// fmt.Printf("x=%v, y=%v\n", x, y)
		// fmt.Println(subGraph)
		// fmt.Println(nodes)
		// fmt.Println(i)
		// fmt.Println()
	}

	fmt.Println(weight)

}

