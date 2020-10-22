package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	br        = bufio.NewReader(os.Stdin)
	bw        = bufio.NewWriter(os.Stdout)
	n, answer int
	c         [][]string
)

//완탐 ㅈㄴ 어렵네
func main() {
	defer bw.Flush()
	fmt.Fscan(br, &n)
	c = make([][]string, n)
	for i := 0; i < n; i++ {
		var str string
		fmt.Fscan(br, &str)
		c[i] = strings.Split(str, "")
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n-1; j++ {
			t := c[i][j]
			c[i][j] = c[i][j+1]
			c[i][j+1] = t
			checkCandy()
			t = c[i][j]
			c[i][j] = c[i][j+1]
			c[i][j+1] = t
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n-1; j++ {
			t := c[j][i]
			c[j][i] = c[j+1][i]
			c[j+1][j] = t
			checkCandy()
			t = c[i][j]
			c[j][i] = c[j+1][i]
			c[j+1][i] = t
		}
	}
	fmt.Fprint(bw, answer)
}
func checkCandy() {
	for i := 0; i < n; i++ {
		cnt := 1
		for j := 0; j < n-1; j++ {
			if c[i][j] == c[i][j+1] {
				cnt++
			} else {
				cnt = 1
			}
			if answer < cnt {
				answer = cnt
			}
		}
	}
	for i := 0; i < n; i++ {
		cnt := 1
		for j := 0; j < n-1; j++ {
			if c[j][i] == c[j+1][i] {
				cnt++
			} else {
				cnt = 1
			}
			if answer < cnt {
				answer = cnt
			}
		}
	}
}
