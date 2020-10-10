package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	br      = bufio.NewReader(os.Stdin)
	bw      = bufio.NewWriter(os.Stdout)
	n       int
	color   [][]byte
	check   [][]bool
	answer1 int
	answer2 int
	dx      = [4]int{0, 0, 1, -1}
	dy      = [4]int{1, -1, 0, 0}
)

func main() {
	fmt.Fscan(br, &n)
	color = make([][]byte, n)
	check = make([][]bool, n)
	for i := 0; i < n; i++ {
		color[i] = make([]byte, n)
		check[i] = make([]bool, n)
		var str string
		fmt.Fscan(br, &str)
		for j := 0; j < n; j++ {
			color[i][j] = str[j]
		}
	}
	find()
	fmt.Print(answer1, answer2)
}
func find() {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if !check[i][j] {
				answer1++
				dfs(i, j, 0)
			}
		}
	}
	for i := range check {
		check[i] = make([]bool, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if !check[i][j] {
				answer2++
				dfs(i, j, 1)
			}
		}
	}
}
func dfs(a, b, choice int) {
	check[a][b] = true
	x, y := 0, 0
	for k := 0; k < 4; k++ {
		x = a + dx[k]
		y = b + dy[k]
		if x >= 0 && y >= 0 && x < n && y < n {
			if !check[x][y] {
				if choice == 0 {
					if color[a][b] == color[x][y] {
						dfs(x, y, 0)
					}
				} else if choice == 1 {
					if color[a][b] == 'R' || color[a][b] == 'G' {
						if color[x][y] == 'R' || color[x][y] == 'G' {
							dfs(x, y, 1)
						}
					} else {
						if color[a][b] == color[x][y] {
							dfs(x, y, 1)
						}
					}
				}
			}
		}
	}
}
