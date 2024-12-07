package main

import (
	"fmt"
	"sort"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MaxProductOfThree(A []int) int {
	sort.Ints(A)

	l := len(A) - 1
	max1 := A[l] * A[l-1] * A[l-2]
	max2 := A[l] * A[0] * A[1]

	return max(max1, max2)
}

func main() {
	fmt.Println("MaxProductOfThree:", MaxProductOfThree([]int{-3, 1, 2, -2, 5, 6}))
	fmt.Println("MaxProductOfThree:", MaxProductOfThree([]int{-20, 1, 2, -2, 5, 6}))
	fmt.Println("MaxProductOfThree:", MaxProductOfThree([]int{20, 1, 2}))
}
