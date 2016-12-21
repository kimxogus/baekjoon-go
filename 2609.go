package main

import "fmt"

func main() {
	var a, b int

	fmt.Scanf("%d %d", &a, &b)

	gcd := getGCD(a, b)

	fmt.Printf("%d\n%d", gcd, getLCM(a, b, gcd))
}

func getGCD(a, b int) int {
	if a < b {
		tmp := a
		a = b
		b = tmp
	}

	for b > 0 {
		tmp := b
		b = a % b
		a = tmp
	}
	return a
}

func getLCM(a, b, gcd int) int {
	a = a / gcd
	b = b / gcd

	return a * b * gcd
}