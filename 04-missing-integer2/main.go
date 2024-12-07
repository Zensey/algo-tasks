package main

import "fmt"

/*
	You are given a list of n-1 integers and these integers are in the range of 1 to n. There are no duplicates in list. One of the integers is missing in the list.

	I/P    [1, 2, 4, ,6, 3, 7, 8]
	O/P    5
*/

func sum(a []int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}

/*
   solution:
   sum=(1+n)*n/2
   missing = sum_calc - sum_act
*/

func missingNumber(in []int) int {
	sumAct := sum(in)
	sum := (1 + len(in) + 1) * (len(in) + 1) / 2
	missing := sum - sumAct
	if missing > 0 {
		fmt.Println("missing number:", missing)
	}
	return missing
}

func main() {
	fmt.Println(missingNumber([]int{1, 2, 4, 6, 3, 7, 8}))
}
