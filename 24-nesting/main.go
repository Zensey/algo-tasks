package main

import (
	"fmt"
)

func Nesting(S string) int {
	open := 0
	for _,i := range S {
		if i == '(' {
			open++
		}
		if i == ')' {
			if open == 0 {
				return 0
			}
			open--
		}
	}
	if open == 0 {
		return 1
	}
	return 0
}

func main() {
	fmt.Println("Nesting:", Nesting("(()()())"))
	fmt.Println("Nesting:", Nesting("(()"))
	fmt.Println("Nesting:", Nesting("(("))
	fmt.Println("Nesting:", Nesting(")))"))
	fmt.Println("Nesting:", Nesting(")("))
	fmt.Println("Nesting:", Nesting("(((((((((((((((((((((((())))))))))))))))))))))))"))
}
