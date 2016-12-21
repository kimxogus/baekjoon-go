package main

import (
	"fmt"
)

var width, height int
var m [][]int

var distance int

func main() {

	fmt.Scanf("%d %d", &height, &width)

	m = make([][]int, height)
	for i := range m {
		m[i] = make([]int, width)

		var row int
		fmt.Scanf("%d", &row)
		for j := range m[i] {
			m[i][width-j-1] = row % 10
			row /= 10
		}
	}

	distance = int(^uint(0) >> 1)
	traverse(0, 0, 1, make(map[int]bool))

	fmt.Printf("%d", distance)
}

func traverse(x, y, d int, v map[int]bool) {
	if x == width-1 && y == height-1 {
		if distance > d {
			distance = d
		}
		return
	}

	v[getKey(x, y)] = true

	if x < width - 1 && hasNotVisited(x+1, y, v) && m[y][x+1] == 1 {
		traverse(x+1, y, d+1, copyMap(v))
	}
	if x > 0 && hasNotVisited(x-1, y, v) && m[y][x-1] == 1 {
		traverse(x-1, y, d+1, copyMap(v))
	}

	if y < height - 1 && hasNotVisited(x, y+1, v) && m[y+1][x] == 1 {
		traverse(x, y+1, d+1, copyMap(v))
	}
	if y > 0 && hasNotVisited(x, y-1, v) && m[y-1][x] == 1 {
		traverse(x, y-1, d+1, copyMap(v))
	}
}

func hasNotVisited(x, y int, v map[int]bool) bool {
	_, hasKey := v[getKey(x, y)]
	return !hasKey
}

func getKey(x, y int) int {
	return x * 10 + y
}

func copyMap(m map[int]bool) map[int]bool {
	n := make(map[int]bool)
	for k, v := range m {
		n[k] = v
	}
	return n
}