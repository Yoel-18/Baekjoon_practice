package main

import (
	"bufio"
	"os"
	"strconv"
)

var br = bufio.NewReader(os.Stdin)
var sc = bufio.NewScanner(os.Stdin)
var bw = bufio.NewWriter(os.Stdout)

var (
	n     int
	g     [100][100]int
	check [100]bool
	q     []int
)

func main() {
	defer bw.Flush()
	n = nextInt()
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			g[i][j] = nextInt()
		}
	}
	find()
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if g[i][j] == 1 {
				bw.WriteString("1 ")
			} else {
				bw.WriteString("0 ")
			}
		}
		bw.WriteString("\n")
	}
}
func find() {
	for i := 0; i < n; i++ {
		check = [100]bool{}
		for j := 0; j < n; j++ {
			if g[i][j] == 1 && !check[j] {
				bfs(i, j)
			}
		}
	}
}
func bfs(x, y int) {
	check[y] = false
	q = append(q, y)
	for len(q) > 0 {
		tmp := q[0]
		q = q[1:len(q)]
		for i := 0; i < n; i++ {
			if !check[i] && g[tmp][i] == 1 {
				q = append(q, i)
				g[x][i] = 1
				check[i] = true
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
