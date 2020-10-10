package main

import (
	"bufio"
	"fmt"
	"os"
)

var br *bufio.Reader = bufio.NewReader(os.Stdin)
var bw *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(bw, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(br, f, a...) }

var (
	t, m, n, k, a, b int
	maap             [50][50]int
	check            [50][50]bool
	count            int
	dx               = [4]int{1, -1, 0, 0}
	dy               = [4]int{0, 0, -1, 1}
)

func main() {
	defer bw.Flush()
	scanf("%d\n", &t)
	for i := 0; i < t; i++ {
		initM()
		scanf("%d %d %d\n", &m, &n, &k)
		for j := 0; j < k; j++ {
			scanf("%d %d\n", &a, &b)
			maap[b][a] = 1
		}
		find()
	}
}
func initM() {
	count = 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			maap[i][j] = 0
			check[i][j] = false
		}
	}
}
func find() {
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if !check[i][j] && maap[i][j] == 1 {
				count++
				dfs(i, j)
			}
		}
	}
	printf("%d\n", count)
}
func dfs(x, y int) {
	check[x][y] = true
	for i := 0; i < 4; i++ {
		xx, yy := x+dx[i], y+dy[i]
		if xx >= 0 && xx < n && yy >= 0 && yy < m {
			if !check[xx][yy] && maap[xx][yy] == 1 {
				dfs(xx, yy)
			}
		}
	}
}
