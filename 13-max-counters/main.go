package main

import (
	"fmt"
)

/*
A small frog wants to get to the other side of a river. The frog is initially located on one bank of the river (position 0) and wants to get to the opposite bank (position X+1). Leaves fall from a tree onto the surface of the river.

You are given an array A consisting of N integers representing the falling leaves. A[K] represents the position where one leaf falls at time K, measured in seconds.

The goal is to find the earliest time when the frog can jump to the other side of the river. The frog can cross only when leaves appear at every position across the river from 1 to X (that is, we want to find the earliest moment when all the positions from 1 to X are covered by leaves). You may assume that the speed of the current in the river is negligibly small, i.e. the leaves do not change their positions once they fall in the river.
*/


func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func MaxCounters(N int, A []int) []int {
	cc := make([]int, N)
	maxV := 0
	lastSet := 0

	for _, va := range A {
		if va <= N {
			n := cc[va-1]
			if n < lastSet {
				n = lastSet
			}
			n = n + 1

			cc[va-1] = n
			maxV = max(maxV, n)

		} else if va == N+1 {
			lastSet = maxV
		}
	}

	// optimisation
	for i := range cc {
		if cc[i] < lastSet {
			cc[i] = lastSet
		}
	}
	return cc
}

func main() {
	fmt.Println("MaxCounters:", MaxCounters(5, []int{3, 4, 4, 6, 1, 4, 4}), "[3 2 2 4 2]")
}
