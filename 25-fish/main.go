package main

import (
	"fmt"
)

func Fish(A, B []int) int {
	s := make([]int, 0) // stack of size value for those fishes who flow upstream

	aliveDownstream := 0
	for i := len(A) - 1; i >= 0; i-- {
		if B[i] == 0 {
			// push
			s = append(s, A[i])
		} else {
			for {
				if len(s) > 0 {
					top := s[len(s)-1]
					if top < A[i] {
						// pop
						s = s[:len(s)-1]
					} else if top > A[i] {
						break
					}
				} else {
					aliveDownstream += 1
					break
				}
			}
		}
	}
	return aliveDownstream + len(s)
}

func main() {
	fmt.Println("Fish:", Fish([]int{4, 3, 2, 1, 5}, []int{0, 1, 0, 0, 0}))
	fmt.Println("Fish:", Fish([]int{4, 3, 10, 1, 5}, []int{0, 1, 0, 0, 0}))
	fmt.Println("Fish:", Fish([]int{4, 3, 10, 1, 5}, []int{0, 1, 1, 0, 0}))
}
