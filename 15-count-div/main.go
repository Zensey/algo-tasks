package main

import (
	"fmt"
)

// naive solution
func CountDiv_(a, b, k int) int {
	count := 0
	for i := a; i <= b; i++ {
		if i%k == 0 {
			// fmt.Println(i)
			count++
		}
	}
	return count
}

func CountDiv(a, b, k int) int {
	count := (b/k - a/k)
	if a%k == 0 {
		count++
		return count
	}
	return count
}

func main() {
	fmt.Println("CountDiv:", CountDiv(10, 20, 10))  //2
	fmt.Println("CountDiv:", CountDiv(6, 11, 2))    //3
	fmt.Println("CountDiv:", CountDiv(11, 345, 17)) //20
	fmt.Println("CountDiv:", CountDiv(10, 10, 5))   //1
	fmt.Println("CountDiv:", CountDiv(10, 10, 7))   //0
	fmt.Println("CountDiv:", CountDiv(10, 10, 20))  //0
	fmt.Println("CountDiv:", CountDiv(10, 10, 2))   //1
}
