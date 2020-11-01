package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var (
	sc          = bufio.NewScanner(os.Stdin)
	tc, n, m, w int
)

func main() {
	tc = nextInt()
	for i := 0; i < tc; i++ {
		n, m, w = nextInt(), nextInt(), nextInt()
		edges := make([]edge, 0)
		for j := 0; j < m; j++ { //	도로 입력, 방향 x
			s, e, t := nextInt(), nextInt(), nextInt()
			edges = append(edges, edge{s, e, t})
			edges = append(edges, edge{e, s, t})
		}
		for j := 0; j < w; j++ { //	웜홀 입력, 방향 o
			s, e, t := nextInt(), nextInt(), nextInt()
			edges = append(edges, edge{s, e, (-1) * t})
		}

		nodes := make([]int, n+1)
		for j := range nodes {
			nodes[j] = math.MaxInt32
		}
		nodes[1] = 0
		isUpdate := false

	loop:
		for j := 1; j <= n; j++ {
			isUpdate = false
			for _, edge := range edges {
				if nodes[edge.end] > nodes[edge.start]+edge.cost {
					nodes[edge.end] = nodes[edge.start] + edge.cost
					isUpdate = true

					if i == n {
						isUpdate = true
						break loop
					}
				}
			}
			if !isUpdate {
				break
			}
		}
		if isUpdate {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

type edge struct {
	start, end, cost int
}

func nextInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}
func init() {
	sc.Split(bufio.ScanWords)
}
