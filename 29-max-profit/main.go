package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func MaxProfit(A []int) int {
	max_profit := 0
	min_day := A[0]

	for i, v := range A {
		if i == 0 {
			continue
		}
		min_day = min(min_day, v)
		max_profit = max(max_profit, v-min_day)
	}
	return max_profit
}

func main() {
	fmt.Println("MaxProfit N:", MaxProfit([]int{23171, 21011, 21123, 21366, 21013, 21367}))
	fmt.Println("MaxProfit N:", MaxProfit([]int{1, 2, 3}))
}
