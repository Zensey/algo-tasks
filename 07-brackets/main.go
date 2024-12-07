package main

import (
	"fmt"
)

// https://www.lavivienpost.com/combinations-of-adding-parentheses

func helper(open, close int, sum *int) {
	if open == 0 && close == 0 {
		*sum += 1
		return
	}
	if open > 0 {
		helper(open-1, close+1, sum)
	}
	if close > 0 {
		helper(open, close-1, sum)
	}
}

func BracketCombinations(num int) int {
	sum := 0
	helper(7, 0, &sum)
	return sum
}

func main() {
	sum := BracketCombinations(7)
	fmt.Println(sum)
}
