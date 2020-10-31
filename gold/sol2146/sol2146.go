package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var (
	sc                    = bufio.NewScanner(os.Stdin)
	n, num, isLand, count int
	min                   = math.MaxInt32
	maap                  [100][100]int
	dx                    = []int{-1, 1, 0, 0}
	dy                    = []int{0, 0, -1, 1}
	check                 [100][100]bool
	isEnd                 bool
)

func main() {
	n = nextInt()
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			maap[i][j] = nextInt()
		}
	}

	//	섬 번호 붙이기
	num = 1
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if !check[i][j] && maap[i][j] == 1 {
				setNumber(i, j, num)
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if maap[i][j] >= 1 {
				num = maap[i][j]
				for d := 0; d < 4; d++ {
					isEnd = false
					if checkRange(i+dx[d], j+dy[d]) {
						continue
					}
					if maap[i+dx[d]][j+dy[d]] == 0 {
						makeBridge(i+dx[d], j+dy[d])
						if min > count {
							min = count
						}
					}
				}
			}
		}
	}
	fmt.Print(min)
}
func setNumber(x, y, num int) {
	q := make([]pos, 0)
	maap[x][y] = num
	q = append(q, pos{x, y})
	check[x][y] = true

	if len(q) > 0 {
		cur := q[0]
		q = q[1:]

		for d := 0; d < 4; d++ {
			nx := cur.x + dx[d]
			ny := cur.y + dy[d]

			if checkRange(nx, ny) {
				continue
			}
			if check[nx][ny] {
				continue
			}
			if maap[nx][ny] == 1 {
				maap[nx][ny] = num
				q = append(q, pos{nx, ny})
				check[nx][ny] = true
			}
		}
	}
}
func makeBridge(x, y int) {
	if checkRange(x, y) {
		return
	}
	for i := 0; i < 4; i++ {
		nx := x + dx[i]
		ny := y + dy[i]

		if !checkRange(nx, ny) && maap[nx][ny] == 0 {
			count++
			makeBridge(nx, ny)
		}
	}

}

func checkRange(x, y int) bool {
	return (x < 0 || x >= n || y < 0 || y >= n)
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
