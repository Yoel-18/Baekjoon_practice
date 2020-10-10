package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	bw = bufio.NewWriter(os.Stdout)
	n  int
	s  [][]bool
)

func main() {
	defer bw.Flush()
	fmt.Scan(&n)
	s = make([][]bool, n)
	for i := 0; i < n; i++ {
		s[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			s[i][j] = false
		}
	}
	star(0, 0, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if s[i][j] {
				bw.WriteByte('*')
			} else {
				bw.WriteByte(' ')
			}
		}
		bw.WriteByte('\n')
	}
}

func star(a, b, size int) {
	if size == 1 {
		s[a][b] = true
		return
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				continue
			}
			star(a+(size/3)*i, b+(size/3)*j, size/3)
		}
	}
}
