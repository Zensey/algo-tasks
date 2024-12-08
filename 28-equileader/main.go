package main

import (
	"fmt"
)

func EquiLeader(A []int) int {
	value := 0
	size := 0

	for _, v := range A {
		if size == 0 {
			value = v
			size += 1
		} else {
			if value == v {
				size += 1
			} else {
				size -= 1
			}
		}
	}
	count := 0
	for _, v := range A {
		if v == value {
			count++
		}
	}

	if count > len(A)/2 {
		countL := 0
		countEL := 0

		for i, v := range A {
			if v == value {
				countL++
			}
			if countL > (i+1)/2 && count-countL > (len(A)-i-1)/2 {
				countEL += 1
			}
		}
		return countEL
	}
	return 0
}


func main() {
	fmt.Println("EquiLeader:", EquiLeader([]int{4, 3, 4, 4, 4, 2}) == 2)
	fmt.Println("EquiLeader:", EquiLeader([]int{3, 3, 1, 3, 1, 3}) == 3)
}
