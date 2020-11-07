package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	br    = bufio.NewReader(os.Stdin)
	bw    = bufio.NewWriter(os.Stdout)
	n, m  int
	maze  [][]byte
	check [50][50][64]bool
	dx    = []int{-1, 1, 0, 0}
	dy    = []int{0, 0, -1, 1}
)

func main() {
	defer bw.Flush()
	fmt.Fscan(br, &n, &m)
	maze = make([][]byte, n)
	startX, startY := 0, 0
	for i := 0; i < n; i++ {
		fmt.Fscan(br, &maze[i])
		for j := 0; j < m; j++ {
			if maze[i][j] == '0' {
				startX, startY = i, j
			}
		}
	}
	fmt.Print(maze)
	findByBFS(startX, startY)
}
func findByBFS(x, y int) {
	q := make([]node, 0)
	q = append(q, node{x, y, 0, 0})
	check[x][y][0] = true

	for len(q) > 0 {
		cur := q[0]
		count := cur.count
		key := cur.key

		if maze[cur.x][cur.y] == '1' {
			fmt.Fprint(bw, cur.count)
			return
		}

		for i := 0; i < 4; i++ {
			nx := cur.x + dx[i]
			ny := cur.y + dy[i]

			if nx < 0 || ny < 0 || nx >= n || ny >= m {
				continue
			}
			if maze[nx][ny] == '#' || check[nx][ny][key] {
				continue
			}
			if maze[nx][ny]-'a' >= 0 && maze[nx][ny]-'a' < 6 {
				curKey := (1 << (maze[nx][ny] - 'a')) | key
				if !check[nx][ny][curKey] {
					check[nx][ny][curKey] = true
					check[nx][ny][key] = true
					q = append(q, node{nx, ny, count + 1, curKey})
				}
			} else if maze[nx][ny]-'A' >= 0 && maze[nx][ny]-'A' < 6 {
				curD := (1 << (maze[nx][ny] - 'A')) & key
				if curD > 0 {
					check[nx][ny][key] = true
					q = append(q, node{nx, ny, count + 1, key})
				}
			}
		}
	}
	fmt.Fprint(bw, -1)
}

type node struct {
	x, y, count, key int
}
