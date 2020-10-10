package main

import (
	"bufio"
	"container/heap"
	"os"
	"strconv"
)

var sc *bufio.Scanner

func nextInt() int {
	sc.Scan()
	b := sc.Bytes()
	ret := 0
	for _, v := range b {
		ret *= 10
		ret += int(v) - 48
	}
	return ret
}

func main() {
	sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	n, m, s := nextInt(), nextInt(), nextInt()
	g := make([][]pair, n+1)
	for i := 0; i < m; i++ {
		a, b, c := nextInt(), nextInt(), nextInt()
		g[a] = append(g[a], pair{b, c})
	}
	q := queue{}
	v := make([]int, n+1)
	heap.Push(&q, &pair{s, 1})
	v[s] = 1

	for q.Len() > 0 {
		here := heap.Pop(&q).(*pair)

		if v[here.x] < here.y {
			continue
		}
		for _, next := range g[here.x] {
			if v[next.x] == 0 || v[next.x] > v[here.x]+next.y {
				v[next.x] = v[here.x] + next.y
				heap.Push(&q, &pair{next.x, v[next.x]})
			}
		}
	}

	for i := 1; i <= n; i++ {
		if v[i] == 0 {
			wr.WriteString("INF")
		} else {
			wr.WriteString(strconv.Itoa(v[i] - 1))
		}
		wr.WriteByte('\n')
	}
}

type pair struct {
	x int
	y int
}

//To implement an interface in Go, we just need to implement all the methods in the interface
type queue []*pair

func (q queue) Len() int {
	return len(q)
}
func (q queue) Less(i, j int) bool {
	return q[i].y < q[j].y
}
func (q queue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}
func (q *queue) Push(x interface{}) {
	*q = append(*q, x.(*pair))
}
func (q *queue) Pop() interface{} {
	old := *q
	n := len(old)
	item := old[n-1]
	*q = old[0 : n-1]
	return item
}
