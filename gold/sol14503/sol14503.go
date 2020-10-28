package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	sc    = bufio.NewScanner(os.Stdin)
	n, m  int
	maap  [50][50]int
	check [50][50]bool
	dPos  = []pos{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
)

//	0 북쪽, 1 동쪽, 2 남쪽, 3 서쪽
func main() {
	n, m = nextInt(), nextInt()
	r, c, d := nextInt(), nextInt(), nextInt()
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			maap[i][j] = nextInt()
		}
	}
	fmt.Print(find(r, c, d))
}
func find(start, end, dir int) int {
	dirCnt := 0
	clean := 0
	nx, ny := 0, 0
	flag := true
	for flag {
		if maap[start][end] == 0 {
			maap[start][end] = 2
			clean++
		}
		for {
			if dirCnt == 4 {
				nx = start - dPos[dir].x
				ny = end - dPos[dir].y
				if maap[nx][ny] == 1 {
					flag = false
					break
				} else {
					start = nx
					end = ny
					dirCnt = 0
				}
			}
			dir = (dir + 3) % 4
			nx = start + dPos[dir].x
			ny = end + dPos[dir].y

			if maap[nx][ny] == 0 {
				dirCnt = 0
				start = nx
				end = ny
			} else {
				dirCnt++
				continue
			}
		}
	}
	return clean
}

type pos struct {
	x, y int
}

func nextInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}
func init() {
	sc.Split(bufio.ScanWords)
}
