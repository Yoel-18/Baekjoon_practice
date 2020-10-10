package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var br = bufio.NewReader(os.Stdin)
var sc = bufio.NewScanner(os.Stdin)
var bw = bufio.NewWriter(os.Stdout)

var (
	w, h  int
	l     [50][50]int
	check [50][50]bool
	dx    = [8]int{1, 1, 0, -1, -1, -1, 0, 1}
	dy    = [8]int{0, 1, 1, 1, 0, -1, -1, -1}
)

func main() {
	for {
		fmt.Scan(&w, &h)
		if w == 0 && h == 0 {
			break
		}
		l = [50][50]int{}
		check = [50][50]bool{}
		for i := 0; i < h; i++ {
			for j := 0; j < w; j++ {
				l[i][j] = nextInt()
			}
		}
		find()
	}
}
func find() {
	count := 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if !check[i][j] && l[i][j] == 1 {
				count++
				dfs(i, j)
			}
		}
	}
	fmt.Println(count)
}
func dfs(x, y int) {
	check[x][y] = true
	for k := 0; k < 8; k++ {
		xx := x + dx[k]
		yy := y + dy[k]
		if xx >= 0 && xx < h && yy >= 0 && yy < w {
			if !check[xx][yy] && l[xx][yy] == 1 {
				dfs(xx, yy)
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
