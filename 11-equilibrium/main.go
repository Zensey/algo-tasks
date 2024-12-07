package main

import (
	"fmt"
)

/*
A non-empty zero-indexed array A consisting of N integers is given. Array A represents numbers on a tape.

Any integer P, such that 0 < P < N, splits this tape into two non-empty parts: A[0], A[1], ..., A[P − 1] and A[P], A[P + 1], ..., A[N − 1].

The difference between the two parts is the value of: |(A[0] + A[1] + ... + A[P − 1]) − (A[P] + A[P + 1] + ... + A[N − 1])|

In other words, it is the absolute difference between the sum of the first part and the sum of the second part.
*/

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func TapeEquilibrium(A []int) int {
	sum := 0
	for _, v := range A {
		sum += v
	}

	suml := 0
	mindif := 0
	for i, v := range A {
		if i == len(A)-1 {
			break
		}
		suml += v
		sumr := sum - suml
		dif := abs(sumr - suml)
		if i == 0 {
			mindif = dif
		} else {
			if dif > mindif { // optimisation
				break
			}
			mindif = min(mindif, dif)
		}
	}
	return mindif
}

func main() {
	fmt.Println("TapeEquilibrium mind:", TapeEquilibrium([]int{3, 1, 2, 4, 3})) // 1
	fmt.Println("TapeEquilibrium mind:", TapeEquilibrium([]int{2, 3, 5, 10}))   // 0
	fmt.Println("TapeEquilibrium mind:", TapeEquilibrium([]int{-1000, 1000}))   // ->2000
}
