package main

import (
	"fmt"
)

func main() {
	var n int
	var h [][]int

	fmt.Scanf("%d", &n)

	// house
	h = make([][]int, n)

	// cost
	c := make([]int, 3)

	for i := 0; i < n; i++ {
		h[i] = make([]int, 3)

		for j := 0; j < 3; j++ {
			fmt.Scanf("%d", &h[i][j])
		}

		for j := 0; j < 3; j++ {
			n1, n2 := getNeighbor(j-1), getNeighbor(j+1)
			if i > 0 {
				c[j] = h[i][j] + min(h[i-1][n1], h[i-1][n2])
			} else {
				c[j] = h[i][j]
			}
			fmt.Printf("%d\t", c[j])
		}
		fmt.Printf("\n")
	}

	fmt.Printf("%d", (c[0] + c[1] + c[2]))
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

func getNeighbor(curr int) int {
	if curr > 2 {
		return 0
	} else if curr < 0 {
		return 2
	}

	return curr
}
