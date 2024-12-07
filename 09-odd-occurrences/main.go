package main

import (
	"fmt"
)



func OddOccurrencesInArray_(A []int) int {
	tab := make(map[int]struct{}, 0)

	for _, v := range A {
		_, ok := tab[v]
		if !ok {
			tab[v] = struct{}{}
		} else {
			delete(tab, v)
		}
	}
	if len(tab) == 1 {
		for k := range tab {
			return k
		}
	}	
	return -1
}

func OddOccurrencesInArray(A []int) int {
	m := 0
	for _, v := range A {
		m = m ^ v
	}
	return m
}

func main() {
	fmt.Println("OddOccurrencesInArray:", OddOccurrencesInArray([]int{9, 3, 9, 3, 9, 7, 9}))
}
