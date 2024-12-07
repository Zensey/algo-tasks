package main

import (
	"fmt"
)

func GenomicRangeQuery(S string, P []int, Q []int) []int {
	r := make([]int, 0)
	pA := make([]int, len(S))
	pC := make([]int, len(S))
	pG := make([]int, len(S))
	pT := make([]int, len(S))

	sumA := 0
	sumC := 0
	sumG := 0
	sumT := 0
	for i := 0; i < len(S); i++ {
		if S[i] == 'A' {
			sumA += 1
		}
		if S[i] == 'C' {
			sumC += 1
		}
		if S[i] == 'G' {
			sumG += 1
		}
		if S[i] == 'T' {
			sumT += 1
		}

		pA[i] = sumA
		pC[i] = sumC
		pG[i] = sumG
		pT[i] = sumT
	}

	calcDif := func(p []int, a, b int) int {
		if a > 0 {
			return p[b] - p[a-1]
		}
		if a == 0 {
			return p[b]
		}
		return 0
	}

	for i := 0; i < len(P); i++ {
		minImpactFactor := 0
		a, b := P[i], Q[i]
		P := [][]int{pA, pC, pG, pT}

		for i, m := range P {
			q := calcDif(m, a, b)
			if q > 0 {
				minImpactFactor = i + 1
				break
			}
		}
		r = append(r, minImpactFactor)
	}
	return r
}

func main() {
	fmt.Println("GenomicRangeQuery:", GenomicRangeQuery("CAGCCTA", []int{2, 5, 0}, []int{4, 5, 6})) // 2 4 1
	fmt.Println("GenomicRangeQuery:", GenomicRangeQuery("A", []int{0}, []int{0}))                   //1
}
