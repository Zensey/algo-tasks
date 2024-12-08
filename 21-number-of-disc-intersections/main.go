package main

import (
	"fmt"
	"sort"
)

// http://www.lucainvernizzi.net/blog/2014/11/21/codility-beta-challenge-number-of-disc-intersections/
// Compute the number of intersections in a sequence of discs.
type Event struct {
	x     int
	begin int
}

type ByX []Event

func (a ByX) Len() int      { return len(a) }
func (a ByX) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByX) Less(i, j int) bool {
	if a[i].x == a[j].x {
		return a[i].begin > a[j].begin
	}
	return a[i].x < a[j].x
}

func NumberOfDiscIntersections(A []int) int {
	ev := make([]Event, 0)
	for c, r := range A {
		ev = append(ev, Event{c - r, 1}, Event{c + r, -1})
	}
	sort.Sort(ByX(ev))

	sum := 0
	activeCircles := 0

	for _, e := range ev {
		if e.begin == 1 {
			if activeCircles > 0 {
				sum += activeCircles
			}
			activeCircles += 1
		} else {
			activeCircles -= 1
		}
	}

	if sum > 10000000 {
		return -1
	}
	return sum
}

func main() {
	fmt.Println("NumberOfDiscIntersections:", NumberOfDiscIntersections([]int{1, 5, 2, 1, 4, 0}))
}
