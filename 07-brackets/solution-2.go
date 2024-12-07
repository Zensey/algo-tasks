package main

import (
	"fmt"
	"math"
)

func fact(x int) int {
	if x == 0 || x == 1 {
		return 1
	}
	return x * fact(x-1)
}

func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func calc_(n int, m map[int]int) int {
	// trivial case
	if n == 1 {
		return 1
	}

	result := 0
	a := make([]int, n)

	sum := func() int {
		sum := 0
		for i := 0; i < n; i++ {
			sum += (i + 1) * a[i]
		}
		return sum
	}
	calc := func(a []int) int {

		// calculate permutations w/ repetitions
		s := 0
		for i := 0; i < n; i++ {
			s += a[i]
		}
		res := fact(s)
		for i := 0; i < n; i++ {
			if a[i] > 0 {
				res = res / fact(a[i])
			}
		}
		fmt.Println("calc > permutations", res, m)

		// multiply res by number of each radix child' combinations
		for i := 1; i < n; i++ {
			if a[i] > 0 {
				res = res * powInt(m[i], a[i]) 
			}
		}
		// fmt.Println("calc >>", res)
		return res
	}

	for {
		a[0] += 1
		s := sum()
		if s == n {
			result += calc(a)
		}
		if s > n {

			// carry into a higher radix
			for i := 1; i < n; i++ {
				a[i] += 1
				a[i-1] = 0
				s := sum()
				if s == n {
					result += calc(a)
					break
				}
				if s < n {
					break
				}
			}

			if a[n-1] == 1 {
				break
			}
		}
	}

	return result
}


// combinatrial solution
func BracketCombinations_(num int) int {
	m := make(map[int]int)

	res := 0
	for i := 1; i <= num; i++ {
		res = calc_(i, m)
		m[i] = res
	}
	return res
}


