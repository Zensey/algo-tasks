package main

import (
	"fmt"
)

func matchRunes(a, b rune) bool {
	if a == '(' && b == ')' {
		return true
	}
	if a == '[' && b == ']' {
		return true
	}
	if a == '{' && b == '}' {
		return true
	}
	return false
}

func Brackets(S string) int {
	s := make([]rune, 0)

	for _, i := range S {
		// fmt.Println(string(rune(i)))

		if i == '(' || i == '[' || i == '{' { // push
			s = append(s, i)
		}
		if i == ')' || i == ']' || i == '}' { // pop
			if len(s) == 0 {
				return 0
			}
			rune := s[len(s)-1]
			if matchRunes(rune, i) { // delete last
				s = s[:len(s)-1]
			} else {
				break
			}
		}
	}
	if len(s) == 0 {
		return 1
	}
	return 0
}

func main() {
	fmt.Println("Brackets:", Brackets("{[()()]}"))
	fmt.Println("Brackets:", Brackets("{[()(]}"))
	fmt.Println("Brackets:", Brackets("{[()"))
	fmt.Println("Brackets:", Brackets("))"))
}
