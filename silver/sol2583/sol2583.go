package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var (
	sc         = bufio.NewScanner(os.Stdin)
	m, n, k, w int
	paper      [100][100]int
	check      [100][100]bool
	answer     = make([]int, 0)
	dx         = [4]int{1, -1, 0, 0}
	dy         = [4]int{0, 0, 1, -1}
)

func main() {
	m, n, k = nextInt(), nextInt(), nextInt()
	for i := 0; i < k; i++ {
		x1, y1, x2, y2 := nextInt(), nextInt(), nextInt(), nextInt()
		for a := x1; a < x2; a++ {
			for b := y1; b < y2; b++ {
				check[a][b] = true
			}
		}
	}
	count := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if !check[i][j] {
				dfs(i, j)
				count++
				answer = append(answer, w)
				w = 0
			}
		}
	}
	sort.Ints(answer)
	fmt.Println(count)
	for i := 0; i < len(answer); i++ {
		fmt.Print(answer[i], " ")
	}
}
func dfs(i, j int) {
	w++
	check[i][j] = true
	for a := 0; a < 4; a++ {
		xx := i + dx[a]
		yy := j + dy[a]
		if xx >= 0 && yy >= 0 && xx < n && yy < m {
			if !check[xx][yy] {
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
