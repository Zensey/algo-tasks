package main

import (
	"fmt"
	"math"
)

func MinPerimeterRectangle(N int) int {
	r := int(math.Sqrt(float64(N)))
	for a := r; a >= 1; a-- {
		if N%a == 0 {
			b := N / a
			return 2 * (a + b)
		}
	}
	return 0
}

func main() {
	fmt.Println(MinPerimeterRectangle(30))
	fmt.Println(MinPerimeterRectangle(32))
	fmt.Println(MinPerimeterRectangle(1))
	fmt.Println(MinPerimeterRectangle(101))
}
