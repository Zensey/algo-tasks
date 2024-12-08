package main

import (
	"fmt"
)


func Dominator(A []int) int {
	cand := 0
	size := 0
	ind := -1

	for i, v := range A {
		if size == 0 {
			cand = v
			ind = i
			size += 1
		} else {
			if cand == v {
				size += 1
			} else {
				size -= 1
			}
		}
	}
	count := 0
	for _, v := range A {
		if v == cand {
			count++
		}
	}
	if count > len(A)/2 {
		return ind
	}
	return -1
}

func main() {
	fmt.Println("Dominator:", Dominator([]int{1, 8, 2, 8, 8}))
	fmt.Println("Dominator:", Dominator([]int{3, 4, 3, 2, 3, -1, 3, 3}))
}
