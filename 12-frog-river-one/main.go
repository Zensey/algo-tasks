package main

import (
	"fmt"
)

/*
A small frog wants to get to the other side of a river. The frog is initially located on one bank of the river (position 0) and wants to get to the opposite bank (position X+1). Leaves fall from a tree onto the surface of the river.

You are given an array A consisting of N integers representing the falling leaves. A[K] represents the position where one leaf falls at time K, measured in seconds.

The goal is to find the earliest time when the frog can jump to the other side of the river. The frog can cross only when leaves appear at every position across the river from 1 to X (that is, we want to find the earliest moment when all the positions from 1 to X are covered by leaves). You may assume that the speed of the current in the river is negligibly small, i.e. the leaves do not change their positions once they fall in the river.
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

func FrogRiverOne(X int, A []int) int {
	sum := 0
	a := make([]int, X+1)

	for i, v := range A {
		if v > len(a)-1 {
			continue
		}

		if a[v] == 0 {
			a[v] = 1
			sum += 1

			if sum == X {
				return i
			}
		}
	}
	fmt.Println(a)
	return -1
}

func main() {
	fmt.Println("FrogRiverOne:", FrogRiverOne(5, []int{1, 3, 1, 4, 2, 3, 5, 4}))
	fmt.Println("FrogRiverOne:", FrogRiverOne(5, []int{1, 3, 1, 4, 2, 3, 4, 5}))
	fmt.Println("FrogRiverOne:", FrogRiverOne(4, []int{1, 3, 1, 4, 2, 3, 4}))
	fmt.Println("FrogRiverOne:", FrogRiverOne(3, []int{1, 3, 1, 2}))
	fmt.Println("FrogRiverOne:", FrogRiverOne(2, []int{1, 1, 2, 2}))
	fmt.Println("FrogRiverOne:", FrogRiverOne(1, []int{1, 1, 1}))
	fmt.Println("FrogRiverOne:", FrogRiverOne(1, []int{1}))
	fmt.Println("FrogRiverOne:", FrogRiverOne(5, []int{3}))
}
