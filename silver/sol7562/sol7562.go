package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	br = bufio.NewReader(os.Stdin)
	bw = bufio.NewWriter(os.Stdout)
	t  int
	dx = []int{-2, -2, -1, -1, 1, 1, 2, 2}
	dy = []int{-1, 1, -2, 2, -2, 2, -1, 1}
)

func main() {
	defer bw.Flush()
	fmt.Fscan(br, &t)
	for k := 0; k < t; k++ {
		var startx, starty, endx, endy, i int
		fmt.Fscan(br, &i)
		fmt.Fscan(br, &startx, &starty)
		fmt.Fscan(br, &endx, &endy)
		fmt.Fprintln(bw, find(startx, starty, endx, endy, i))
	}
}
func find(x1, y1, x2, y2, ii int) int {
	s := make([][]int, ii)
	for i := range s {
		s[i] = make([]int, ii)
	}
	count := 0
	q := []pos{{x1, y1}}
	s[x1][y1] = 1
	for len(q) > 0 {
		o := q[0]
		if o.x == x2 && o.y == y2 {
			count = s[o.x][o.y] - 1
			break
		}
		q = q[1:]
		for k := 0; k < 8; k++ {
			x, y := o.x+dx[k], o.y+dy[k]
			if 0 <= x && x < ii && 0 <= y && y < ii && s[x][y] == 0 {
				s[x][y] = s[o.x][o.y] + 1
				q = append(q, pos{x, y})
			}
		}
	}
	return count
}

type pos struct {
	x, y int
}
