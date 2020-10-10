package main

import (
	"bufio"
	"os"
	"strconv"
)

var br = bufio.NewReader(os.Stdin)
var sc = bufio.NewScanner(os.Stdin)
var bw = bufio.NewWriter(os.Stdout)

var (
	tomato  [100][100][100]int
	n, m, h int
	dx      = [6]int{0, 0, 1, -1, 0, 0}
	dy      = [6]int{0, 0, 0, 0, 1, -1}
	dz      = [6]int{1, -1, 0, 0, 0, 0}
	q       []info
)

func main() {
	n, m, h := nextInt(), nextInt(), nextInt()
	for k := 0; k < h; k++ {
		for i := 0; i < m; i++ {
			for j := 0; j < n; k++ {
				tomato[k][m][n] = nextInt()
				if tomato[k][m][n] == 1 {
					q = append(q, info{})
				}
			}
		}
	}
}

type info struct {
	x   int
	y   int
	z   int
	day int
}

func nextInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}
func init() {
	sc.Split(bufio.ScanWords)
}
