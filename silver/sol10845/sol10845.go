package main

import "fmt"

func main() {
	q := make([]int, 0)
	push(&q, 1)
	push(&q, 2)
	push(&q, 3)
	push(&q, 4)
	push(&q, 5)
	push(&q, 6)
	pop(&q)
	fmt.Println(q, size(q), empty(q), front(q), back(q))
}
func push(q *[]int, x int) {
	*q = append(*q, x)
}
func pop(q *[]int) int {
	nq := *q
	x := nq[0]
	*q = nq[1:]
	return x
}
func size(q []int) int {
	return len(q)
}
func empty(q []int) int {
	if len(q) == 0 {
		return 1
	}
	return 0
}
func front(q []int) int {
	if size(q) == 0 {
		return -1
	}
	return q[0]
}
func back(q []int) int {
	if size(q) == 0 {
		return -1
	}
	return q[size(q)-1]
}
