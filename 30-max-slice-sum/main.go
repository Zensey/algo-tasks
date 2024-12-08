package main

import (
	"fmt"
)

// Find a maximum sum of a compact subsequence of array elements.

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MaxSliceSum(A []int) int {
	maxEnding := 0
	maxSlice := 0

	for _, a := range A {
		maxEnding = max(0, maxEnding+a)
		maxSlice = max(maxSlice, maxEnding)

		fmt.Println(a, maxEnding, maxSlice)
	}

	return maxSlice
}

func main() {
	fmt.Println("MaxSliceSum:", MaxSliceSum([]int{3, 2, -6, 4, 0}))
	fmt.Println("MaxSliceSum:", MaxSliceSum([]int{3, 2, -6, 7, 0}))
	// fmt.Println("MaxSliceSum:", MaxSliceSum([]int{5, 5, -6, 7, 0}))
	// fmt.Println("MaxSliceSum:", MaxSliceSum([]int{-10}))
	// fmt.Println("MaxSliceSum:", MaxSliceSum([]int{-10, -3, -2}))
}
