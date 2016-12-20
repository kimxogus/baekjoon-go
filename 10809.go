package main

import "fmt"

func main() {
	var s string
	fmt.Scanf("%s", &s)

	alphabets := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}

	size := len(s)
	for _, c := range alphabets {
		for i := 0; i < size; i++ {
			if rune(s[i]) == c {
				fmt.Printf("%d ", i)
				break
			}
			if i == size - 1 {
				fmt.Printf("%d ", -1)
			}
		}
	}
}
