package main

import (
	"fmt"
)

var length int
var heap []int

func main() {
	fmt.Scanf("%d", &length)

	heap = make([]int, length)

	var input int
	for i := range heap {
		fmt.Scanf("%d", &input)
		insert(input, i)
	}

	for range heap {
		fmt.Printf("%d\n", pop())
	}
}

func pop() (head int) {
	head = heap[0]

	length--
	if length == 0 {
		return
	}

	heap[0] = heap[length]

	curr := 0
	for {
		min := curr
		left := toLeftChild(curr)
		right := toRightChild(curr)

		if left < length && heap[min] > heap[left] {
			min = left
		}
		if right < length && heap[min] > heap[right] {
			min = right
		}

		if min == curr {
			break
		}

		swap(curr, min)
		curr = min
	}

	return
}

func insert(input, size int) {
	heap[size] = input

	curr := size
	for curr > 0 {
		parent := toParent(curr)

		if heap[parent] > heap[curr] {
			swap(parent, curr)
			if parent > 0 {
				curr = parent
				continue
			}
		}

		break
	}
}

func swap(a, b int) {
	tmp := heap[a]
	heap[a] = heap[b]
	heap[b] = tmp
}

func toParent(i int) int {
	return (i + 1) / 2 - 1
}

func toLeftChild(i int) int {
	return (i + 1) * 2 - 1
}

func toRightChild(i int) int {
	return (i + 1) * 2
}