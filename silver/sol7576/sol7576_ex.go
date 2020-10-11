package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

var (
	tomato      [1000][1000]int
	check       [1000][1000]bool
	n, m, count int
	dx          = [4]int{1, 0, -1, 0}
	dy          = [4]int{0, 1, 0, -1}
	q           []pos
)

func main() {
	m, n = nextInt(), nextInt()
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			tomato[i][j] = nextInt()
			if tomato[i][j] == 1 {
				q = append(q, pos{i, j})
			}
		}
	}
	find()
	isAll()
}
func find() {
	for len(q) > 0 {
		xx := q[0].x
		yy := q[0].y
		q = q[1:]
		check[xx][yy] = true
		for i := 0; i < 4; i++ {
			xxx := xx + dx[i]
			yyy := yy + dy[i]
			if xxx >= 0 && yyy >= 0 && xxx < n && yyy < m {
				if !check[xxx][yyy] && tomato[xxx][yyy] == 0 {
					q = append(q, pos{xxx, yyy})
					check[xxx][yyy] = true
					tomato[xxx][yyy] = tomato[xx][yy] + 1
					count = tomato[xxx][yyy]
				}
			}
		}
	}
}
func isAll() {
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if tomato[i][j] == 0 {
				fmt.Print(-1)
				return
			}
		}
	}
	fmt.Print(count - 1)
}

type pos struct {
	x int
	y int
}

func nextInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}
func init() {
	sc.Split(bufio.ScanWords)
}
