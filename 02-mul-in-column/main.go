package main

import "fmt"

func sum(a, b []int) []int {
	res := make([]int, 0)
	maxlen := len(a)
	if len(b) > maxlen {
		maxlen = len(b)
	}

	carry := 0
	for i := 0; i < maxlen; i++ {
		r1 := 0
		r2 := 0
		if i < len(a) {
			r1 = a[i]
		}
		if i < len(b) {
			r2 = b[i]
		}

		s := r1 + r2 + carry
		if s > 9 {
			carry = 1
			s = s % 10
		}
		res = append(res, s)
	}
	if carry > 0 {
		res = append(res, carry)
	}

	//reverse
	res_ := make([]int, 0)
	for i := len(res) - 1; i >= 0; i-- {
		res_ = append(res_, res[i])
	}
	return res_
}

func mul(a, b []int) []int {
	res := make([]int, 0)
	res = append(res, 0)

	for i := 0; i < len(b); i++ {

		part := make([]int, 0)
		carry := 0
		for j := 0; j < len(a); j++ {
			s := b[i]*a[j] + carry
			if s > 9 {
				carry = s / 10
				s = s % 10
			}
			part = append(part, s)
		}
		if carry > 0 {
			part = append(part, carry)
		}
		// shift part i times
		for k := 0; k < i; k++ {
			part = append(part, 0)
		}

		//reverse
		part_ := make([]int, 0)
		for i := len(part) - 1; i >= 0; i-- {
			part_ = append(part_, part[i])
		}

		res = sum(res, part_)
	}
	return res
}

func main() {
	// a := []int{0, 3, 3, 3}
	// b := []int{3, 3, 3}
	// res := sum(a, b)

	a := []int{1, 1, 1}
	b := []int{3, 3}
	res := mul(a, b)

	fmt.Println(res)
}
