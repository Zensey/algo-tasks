package main

import "fmt"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func findFirstPositive(a []int) (i int, v int) {
	for i := 0; i < len(a); i++ {
		if a[i] > 0 {
			return i, a[i]
		}
	}
	return 0, 0
}

/*
	Find the smallest positive number missing from an unsorted array
	Input: nums = [3,4,-1,1]
	Output: 2
	Example 3:
	Input: nums = [7,8,9,11,12]
	Output: 1
*/

func main() {
	a := []int{2, 3, -7, 6, 8, 1, -10, 15}

	_, first := findFirstPositive(a)
	for i := range a {
		if a[i] <= 0 {
			a[i] = first
		}
	}
	for i := range a {
		j := abs(a[i]) - 1
		if j < len(a) && j >= 0 && a[j] > 0 {
			a[j] = -a[j]
		}
	}
	firstI, _ := findFirstPositive(a)
	fmt.Println(firstI + 1)
}
