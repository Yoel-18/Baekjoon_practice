package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	br                               = bufio.NewReader(os.Stdin)
	bw                               = bufio.NewWriter(os.Stdout)
	n, m, count                      int
	board                            [][]string
	redX, redY, blueX, blueY, hX, hY int
	dPos                             = []point{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
)

func main() {
	defer bw.Flush()
	fmt.Fscan(br, &n, &m)
	board = make([][]string, n)
	for i := 0; i < n; i++ {
		board[i] = make([]string, m)
		var str string
		fmt.Fscan(br, &str)
		board[i] = strings.Split(str, "")
		for j := 0; j < m; j++ {
			if board[i][j] == "R" {
				redX, redY = i, j
			}
			if board[i][j] == "B" {
				blueX, blueY = i, j
			}
			if board[i][j] == "O" {
				hX, hY = i, j
			}
		}
	}
	find()
	fmt.Fprint(bw, count)
}
func find() {
	redQ := make([]point, 0)
	blueQ := make([]point, 0)
}

type point struct {
	x, y int
}
