package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	sc     = bufio.NewScanner(os.Stdin)
	dx     = [4]int{1, -1, 0, 0}
	dy     = [4]int{0, 0, 1, -1}
	maap   [100][100]int
	check  [100][100]bool
	n, max int
)

func main() {
	n = nextInt()
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			maap[i][j] = nextInt()
			if max < maap[i][j] {
				max = maap[i][j]
			}
		}
	}
	answer := 1
	for i := 1; i < max; i++ {
		check = [100][100]bool{}
		cnt := find(i)
		if answer < cnt {
			answer = cnt
		}
	}
	fmt.Print(answer)
}
func find(h int) int {
	count := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if !check[i][j] && maap[i][j] > h {
				dfs(i, j, h)
				count++
			}
		}
	}
	return count
}
func dfs(a, b, h int) {
	check[a][b] = true
	for i := 0; i < 4; i++ {
		xx := a + dx[i]
		yy := b + dy[i]
		if xx >= 0 && yy >= 0 && xx < n && yy < n {
			if !check[xx][yy] && maap[xx][yy] > h {
				dfs(xx, yy, h)
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
