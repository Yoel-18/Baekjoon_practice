package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var (
	sc        = bufio.NewScanner(os.Stdin)
	n, k, sum int
	c         []int
	js        []jewelry
)

func main() {
	n, k = nextInt(), nextInt()
	c = make([]int, k)
	jpq := new(jHeap)
	for i := 0; i < n; i++ {
		a, b := nextInt(), nextInt()
		js = append(js, jewelry{a, b})
	}
	for i := 0; i < k; i++ {
		c[i] = nextInt()
	}
	sort.Slice(js, func(i, j int) bool {
		return js[i].m < js[j].m
	})
	sort.Ints(c)

	o := 0
	for i := 0; i < k; i++ {
		for o < n {
			if c[i] >= js[o].m {
				heap.Push(jpq, js[o])
				o++
			} else {
				break
			}
		}
		if jpq.Len() > 0 {
			cur := heap.Pop(jpq).(jewelry)
			sum += cur.v
		}
	}
	fmt.Print(sum)
}

type jewelry struct {
	m, v int
}
type jHeap []jewelry

func (h jHeap) Len() int {
	return len(h)
}
func (h jHeap) Less(i, j int) bool {
	return h[i].v > h[j].v
}
func (h jHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *jHeap) Push(x interface{}) {
	*h = append(*h, x.(jewelry))
}
func (h *jHeap) Pop() interface{} {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}
func nextInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}
func init() {
	sc.Split(bufio.ScanWords)
}
