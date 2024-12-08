package main

import (
	"fmt"
	"sort"
)

// Determine whether a given string of parentheses (multiple types) is properly nested.
func triangular(p, q, r int) bool {
	return p+q > r && p+r > q && r+q > p
}

func Triangle(A []int) int {
	sort.Sort(sort.IntSlice(A))
	for i := 0; i < len(A)-2; i++ {
		if triangular(A[i], A[i+1], A[i+2]) {
			return 1
		}
	}
	return 0
}

func main() {
	fmt.Println("Triangle:", Triangle([]int{10, 2, 5, 1, 8, 20}))
	fmt.Println("Triangle:", Triangle([]int{10, 50, 5, 1}))
}
