package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

var (
	br        = bufio.NewReader(os.Stdin)
	bw        = bufio.NewWriter(os.Stdout)
	n, answer int
)

func main() {
	defer bw.Flush()
	fmt.Fscan(br, &n)
	num := new(intHeap)
	for i := 0; i < n; i++ {
		var number int
		fmt.Fscan(br, &number)
		heap.Push(num, number)
	}
	for num.Len() > 1 {
		a := heap.Pop(num).(int)
		b := heap.Pop(num).(int)
		answer += a + b
		heap.Push(num, (a + b))
	}
	fmt.Fprint(bw, answer)
}

type intHeap []int

func (h intHeap) Len() int {
	return len(h)
}
func (h intHeap) Less(i, j int) bool {
	return h[i] < h[j]
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
