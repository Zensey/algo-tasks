package main

import (
	"fmt"
	"sort"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

/*
naive approach
*/
func PermCheck_(A []int) int {
	sort.Ints(A)
	for i, v := range A {
		if i+1 != v {
			return 0
		}
	}
	return 1
}

func PermCheck(A []int) int {
	n := len(A)
	for _, v := range A {

		v = abs(v)
		if v > n {
			return 0
		}

		if A[v-1] < 0 {
			return 0
		}
		A[v-1] = -A[v-1]
	}
	return 1
}

func main() {
	fmt.Println("PermCheck:", PermCheck([]int{1, 3, 4, 2}))
	fmt.Println("PermCheck:", PermCheck([]int{1, 3, 4}))
	fmt.Println("PermCheck:", PermCheck([]int{1, 2, 3, 4}))
	fmt.Println("PermCheck:", PermCheck([]int{1, 2, 3, 5}))
}
