package main

import (
	"fmt"
	"sort"
)

func Distinct(A []int) int {
	sort.Ints(A)
	count := 1
	for i := 1; i < len(A); i++ {
		if A[i-1] != A[i] {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println("Distinct:", Distinct([]int{0, 1, 1, 1, 2, 2, 0}))
}
