package main

import "fmt"

func main() {
	var t int
	fmt.Scanf("%d", &t)

	var n int
	r := make([]int, 5, 100)
	r[0] = 1
	r[1] = 1
	r[2] = 1
	r[3] = 2
	r[4] = 2
	for i := 0; i < t; i++ {
		fmt.Scanf("%d", &n)

		for j := len(r); j < n; j++ {
			r = append(r, r[j-1] + r[j-5])
		}

		fmt.Printf("%d\n", r[n-1])
	}
}
