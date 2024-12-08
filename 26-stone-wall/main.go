package main

import (
	"fmt"
)

func StoneWall(H []int) int {
	S := make([]int, 0)
	count := 0

	for _, h := range H {
		for {
			if len(S) > 0 {
				// pop all values on stack > h
				top := S[len(S)-1]
				if h >= top {
					if h > top {
						S = append(S, h)
					}
					break
				}
				S = S[:len(S)-1]
				count++
			}
			if len(S) == 0 {
				S = append(S, h)
				break
			}
		}
	}
	count += len(S)
	return count
}

func main() {
	fmt.Println("StoneWall:", StoneWall([]int{8, 8, 5, 7, 9, 8, 7, 4, 8})) //7
	fmt.Println("StoneWall:", StoneWall([]int{1, 8, 1, 8, 1}))             //3
	fmt.Println("StoneWall:", StoneWall([]int{1, 8, 2, 8, 1}))             //4
}
