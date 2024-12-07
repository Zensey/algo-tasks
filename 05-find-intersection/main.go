package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func FindIntersection(strArr []string) string {

	strToArrayInt := func(str string) []int {
		as := strings.Split(str, ",")
		a := make([]int, 0)
		for _, v := range as {
			i, _ := strconv.Atoi(strings.TrimSpace(v))
			a = append(a, i)
		}
		sort.Ints(a)
		return a
	}

	a := strToArrayInt(strArr[0])
	b := strToArrayInt(strArr[1])
	fmt.Println(a, b)

	I := make([]int, 0)

	ia, ib := 0, 0
	for {
		if ia == len(a) {
			break
		}
		if ib == len(b) {
			break
		}

		if a[ia] == b[ib] {
			I = append(I, a[ia])
			ia++
			ib++
		} else if a[ia] > b[ib] {
			ib++
		} else if a[ia] < b[ib] {
			ia++
		}

	}

	if len(I) == 0 {
		return "false"
	}
	s := ""
	for _, v := range I {
		if s == "" {
			s = strconv.Itoa(v)
		} else {
			s += "," + strconv.Itoa(v)
		}
	}
	return s
}

func main() {
	fmt.Println(FindIntersection([]string{"1, 3, 4, 7, 13", "1, 2, 4, 13, 15"}))
	fmt.Println(FindIntersection([]string{"1, 3, 9, 10, 17, 18", "1, 4, 9, 10"}))
}
