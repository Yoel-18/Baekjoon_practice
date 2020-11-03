package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	sc       = bufio.NewScanner(os.Stdin)
	n, m     int
	g        [][]int
	indegree []int
)

func main() {
	n, m = nextInt(), nextInt()
	g = make([][]int, n+1)
	indegree = make([]int, n+1)
	for i := 1; i <= n; i++ {
		g[i] = make([]int, n+1)
	}
	for i := 0; i < m; i++ {
		a, b := nextInt(), nextInt()
		g[a] = append(g[a], b)
		indegree[b]++
	}

	q := make([]int, 0)
	for i := 1; i <= n; i++ {
		if indegree[i] == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		fmt.Print(q[0], " ")
		cur := q[0]
		q = q[1:]

		for i := 0; i < len(g[cur]); i++ {
			next := g[cur][i]
			indegree[next]--
			if indegree[next] == 0 {
				q = append(q, next)
			}
		}
	}

}
func nextInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}
func init() {
	sc.Split(bufio.ScanWords)
}
