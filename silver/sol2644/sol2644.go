package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	sc                   = bufio.NewScanner(os.Stdin)
	n, m, countX, countY int
	family               [101][101]int
	check                [101]int
)

func main() {
	n = nextInt()
	countX, countY = nextInt(), nextInt()
	m = nextInt()
	for i := 0; i < m; i++ {
		x, y := nextInt(), nextInt()
		family[x][y], family[y][x] = 1, 1
	}
	check[countX] = 1
	for i := 1; i <= n; i++ {
		if family[countX][i] == 1 && check[i] == 0 {
			findByDFS(i, countY, 1)
		}
	}
	fmt.Print(check[countY] - 1)
}
func findByDFS(x, countY, depth int) {
	depth++
	check[x] = depth
	if x == countY {
		return
	}
	for i := 1; i <= n; i++ {
		if family[x][i] == 1 && check[i] == 0 {
			findByDFS(i, countY, depth)
		}
	}
}
func findByBFS() {
	q := make([]int, 0)
	q = append(q, countX)
	check[countX] = 1
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		if node == countY {
			break
		}
		for i := 1; i <= n; i++ {
			if family[node][i] == 1 && check[i] == 0 {
				q = append(q, i)
				check[i] = check[node] + 1
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
