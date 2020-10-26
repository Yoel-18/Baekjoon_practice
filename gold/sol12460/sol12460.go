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
	check                            [10][10][10][10]bool
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
	if find() && count <= 10 {
		fmt.Fprint(bw, count)
	} else {
		fmt.Fprint(bw, -1)
	}
}
func find() bool {
	q := make([][5]int, 0)
	q = append(q, [5]int{redX, redY, blueX, blueY, 0})
	check[redX][redY][blueX][blueY] = true
	for len(q) > 0 {
		rx, ry, bx, by, dis := q[0][0], q[0][1], q[0][2], q[0][3], q[0][4]
		q = q[1:]
		for i := 0; i < 4; i++ {
			newRX, newRY, newBX, newBY, newDis := rx, ry, bx, by, dis+1
			countR, countB := 0, 0
			for board[newRX+dPos[i].x][newRY+dPos[i].y] != "#" && board[newRX][newRY] != "O" {
				newRX += dPos[i].x
				newRY += dPos[i].y
				countR++
			}
			for board[newBX+dPos[i].x][newBY+dPos[i].y] != "#" && board[newBX][newBY] != "O" {
				newBX += dPos[i].x
				newBY += dPos[i].y
				countB++
			}
			if newBX == hX && newBY == hY {
				continue
			}
			if newRX == hX && newRY == hY {
				count = newDis
				return true
			}
			if newRX == newBX && newRY == newBY {
				if countR > countB {
					newRX -= dPos[i].x
					newRY -= dPos[i].y
				} else {
					newBX -= dPos[i].x
					newBY -= dPos[i].y
				}
			}
			if !check[newRX][newRY][newBX][newBY] {
				check[newRX][newRY][newBX][newBY] = true
				q = append(q, [5]int{newRX, newRY, newBX, newBY, newDis})
			}
		}
	}
	return false
}

type point struct {
	x, y int
}
