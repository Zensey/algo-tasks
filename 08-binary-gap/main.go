package main

import (
	"fmt"
)

func maxGap(N int) int {
	s := fmt.Sprintf("%b", N)
	prev := -1
	max := 0
	for i, j := range s {
		if j == '1' {
			if prev >= 0 {
				dif := i - prev - 1
				if dif > max {
					max = dif
				}
			}
			prev = i
		}
	}
	return max
}

func main() {
	fmt.Println("maxGap>", maxGap(5123)) // 1010000000011
	fmt.Println("maxGap>", maxGap(1041)) // 5
	fmt.Println("maxGap>", maxGap(32))   // 0
	fmt.Println("maxGap>", maxGap(1))    // 0
	fmt.Println("maxGap>", maxGap(0))    // 0
}
