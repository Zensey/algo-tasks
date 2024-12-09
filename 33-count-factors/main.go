package main

import (
	"fmt"
)

func CountFactors(n int) int {
	i := 1
	result := 0
	for i*i < n {
		if n%i == 0 {
			result += 2
		}
		i += 1
	}
	if i*i == n {
		result += 1
	}
	return result
}

func main() {
	fmt.Println(CountFactors(24))
}
