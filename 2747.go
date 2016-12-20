package main

import "fmt"

func main() {
	var f []int64
	var size int

	fmt.Scanf("%d", &size)

	if size < 2 {
		fmt.Printf("%d", size)
		return
	}

	f = make([]int64, size)

	f[0] = 0
	f[1] = 1

	for i := 2; i < size; i++ {
		f[i] = f[i - 1] + f[i - 2]
	}

	fmt.Printf("%d", f[size - 1] + f[size - 2])
}