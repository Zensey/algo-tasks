package main

import (
	"fmt"
)

// a solution using running sum

func MinAvgTwoSlice(A []int) int {
	sum2 := 0
	sum3 := 0
	avgMin := float32(0)
	iSlice := -1

	for i := range A {
		sum2 += A[i]
		sum3 += A[i]
		if i > 0 {
			avg := float32(sum2) / 2
			if iSlice != -1 {
				if avg < avgMin {
					iSlice = i - 1
					avgMin = avg
				}
			} else {
				iSlice = 0
				avgMin = avg
			}
			sum2 -= A[i-1]
		}
		if i > 1 {
			avg := float32(sum3) / 3
			if avg < avgMin {
				iSlice = i - 2
				avgMin = avg
			}
			sum3 -= A[i-2]
		}
	}
	return iSlice
}

func main() {
	fmt.Println("MinAvgTwoSlice:", MinAvgTwoSlice([]int{4, 2, 2, 5, 1, 5, 8}))
	fmt.Println("MinAvgTwoSlice:", MinAvgTwoSlice([]int{3, 0, 0, 2, 1}))
	fmt.Println("MinAvgTwoSlice:", MinAvgTwoSlice([]int{1, 3, 1, 2, 2}))
	fmt.Println("MinAvgTwoSlice:", MinAvgTwoSlice([]int{9, 0, 0, 9, 1, 1, 1}))
	fmt.Println("MinAvgTwoSlice:", MinAvgTwoSlice([]int{1, 1, 1, -1, -1, -1, -1, -1}))
}
