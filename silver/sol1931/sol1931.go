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
	n, count int
	c        []con
)

func main() {
	defer bw.Flush()
	fmt.Fscan(br, &n)
	for i := 0; i < n; i++ {
		var xx, yy int
		fmt.Fscan(br, &xx, &yy)
		c = append(c, con{xx, yy})
	}
	sort.Sort(conList(c))
	end := -1
	for i := 0; i < n; i++ {
		if c[i].x >= end {
			end = c[i].y
			count++
		}
	}
	fmt.Fprint(bw, count)
}

type con struct {
	x int
	y int
}
type conList []con

func (c conList) Len() int {
	return len(c)
}
func (c conList) Less(i, j int) bool {
	if c[i].y == c[j].y {
		return c[i].x < c[j].x
	}
	return c[i].y < c[j].y
}
func (c conList) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
