package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	br       = bufio.NewReader(os.Stdin)
	bw       = bufio.NewWriter(os.Stdout)
	maap     []string
	n, count int
	check    [][]bool
	dx       = []int{1, -1, 0, 0}
	dy       = []int{0, 0, 1, -1}
	answer   = make([]int, 0)
)

func main() {
	defer bw.Flush()

	fmt.Fscanln(br, &n)
	maap = make([]string, 0)
	check = make([][]bool, n)

	//	입력
	for i := 0; i < n; i++ {
		check[i] = make([]bool, n)
		cur, _ := br.ReadString('\n')
		maap = append(maap, cur)
	}

	//	탐색
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if maap[i][j] == '1' && !check[i][j] {
				count = 1
				answer = append(answer, find(i, j))
			}
		}

	}
	sort.Ints(answer)

	//	정답출력
	fmt.Fprintln(bw, len(answer))
	for _, num := range answer {
		fmt.Fprintln(bw, num)
	}
}

func find(x, y int) int {
	check[x][y] = true

	for i := 0; i < 4; i++ {
		nx, ny := x+dx[i], y+dy[i]

		if nx >= 0 && ny >= 0 && nx < n && ny < n {
			if maap[nx][ny] == '1' && !check[nx][ny] {
				find(nx, ny)
				count++
			}
		}
	}

	return count
}
