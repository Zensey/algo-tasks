package main

import (
	"fmt"
)

// a solution using running sum

func PassingCars(A []int) int {
	sum0 := 0  // running sum of zero-elements
	count := 0 // count pass

	for _, v := range A {
		if v == 0 {
			sum0 += 1
		}
		if v == 1 {
			count += sum0
		}
	}
	if count > 1000000000 {
		return -1
	}
	return count
}

func main() {
	fmt.Println("PassingCars:", PassingCars([]int{0, 1, 0, 1, 1}))
}
