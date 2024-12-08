package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func RectanglesStrip(A []int, B []int) int {
	f := make(map[int]int)

	for _, v := range A {
		f[v] = f[v] + 1
	}
	for _, v := range B {
		f[v] = f[v] + 1
	}

	maxF := 0
	maxFVal := 0
	for i, v := range f {
		if v > maxF {
			maxF = v
			maxFVal = i
		}
	}

	match := 0
	for i := range A {
		if A[i] == maxFVal || B[i] == maxFVal {
			match += 1
		}
	}
	return match
}

func main() {
	fmt.Println(RectanglesStrip([]int{2, 3, 2, 3, 5}, []int{3, 4, 2, 4, 2}))
	fmt.Println(RectanglesStrip([]int{2, 3, 1, 3}, []int{2, 3, 1, 3}))
	fmt.Println(RectanglesStrip([]int{2, 10, 4, 1, 4}, []int{4, 1, 2, 2, 5}))
}
