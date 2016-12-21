package main

import "fmt"

func main() {
	l := 64
	v := 64

	var t int
	fmt.Scanf("%d", &t)

	var i int
	for i = 1; v != t && l > 0; i++ {
		l /= 2
		if v - l >= t {
			v -= l
			i--
		}
	}

	fmt.Printf("%d", i)
}
