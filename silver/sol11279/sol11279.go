package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	var n int
	br := bufio.NewReader(os.Stdin)
	bw := bufio.NewWriter(os.Stdout)
	defer bw.Flush()
	fmt.Fscan(br, &n)
	h := new(intHeap)
	for i := 0; i < n; i++ {
		var x int
		fmt.Fscan(br, &x)
		if x == 0 {
			v := 0
			if h.Len() > 0 {
				v = heap.Pop(h).(int)
			} else {
				v = 0
			}
			fmt.Fprintln(bw, v)
		} else {
			heap.Push(h, x)
		}
	}
}

type intHeap []int

func (h intHeap) Len() int {
	return len(h)
}
func (h intHeap) Less(i, j int) bool {
	return h[i] > h[j]
}
func (h intHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *intHeap) Pop() interface{} {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}
