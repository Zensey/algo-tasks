package main

/*
Given this linked list: 1->2->3->4->5, and an integer k
    For k = 2, you should return: 2->1->4->3->5
    For k = 3, you should return: 3->2->1->4->5
*/
func main() {
	l := List{}
	l.push(5)
	l.push(4)
	l.push(3)
	l.push(2)
	l.push(1)

	l.head.print()
	new := l.head.reverseN(2)
	new.print()

	return
}
