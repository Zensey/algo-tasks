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

func MaxDoubleSliceSum(A []int) int {
	L := make([]int, len(A))
	R := make([]int, len(A))

	for i := 1; i < len(A)-1; i++ {
		L[i] = max(0, L[i-1]+A[i])
	}
	for i := len(A) - 2; i > 0; i-- {
		R[i] = max(0, R[i+1]+A[i])
	}
	maxS := 0
	for i := 1; i < len(A)-1; i++ {
		maxS = max(maxS, L[i-1]+R[i+1])
	}
	return maxS

}

func main() {
	fmt.Println("MaxDoubleSliceSum:", MaxDoubleSliceSum([]int{1, 2, -1, -1, 2, 1}))
	fmt.Println("MaxDoubleSliceSum:", MaxDoubleSliceSum([]int{5, 17, 0, 3}))
	fmt.Println("MaxDoubleSliceSum:", MaxDoubleSliceSum([]int{3, 2, 6, 4, 5, -15, 10, 2}))
	fmt.Println("MaxDoubleSliceSum:", MaxDoubleSliceSum([]int{6, 1, 5, 6, 4, 2, 9, 4}))
}
