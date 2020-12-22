package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	br = bufio.NewReader(os.Stdin)
	bw = bufio.NewWriter(os.Stdout)
)

func main() {
	defer bw.Flush()
	var m, n int
	fmt.Fscanf(br, "%d %d", &n, &m)
	c := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(br, &c[i])
	}
	min := 80
	for i := 0; i < n-7; i++ {
		for j := 0; j < m-7; j++ {
			countB, countW := 0, 0
			for k := i; k < i+8; k++ {
				for t := j; t < j+8; t++ {
					if (k+t)%2 == 0 {
						if c[k][t] == 'B' {
							countW++
						} else {
							countB++
						}
					} else {
						if c[k][t] == 'B' {
							countB++
						} else {
							countW++
						}
					}
				}
			}
			min = minInt(min, countB)
			min = minInt(min, countW)
		}
	}
	fmt.Fprint(bw, min)
}
func minInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}
