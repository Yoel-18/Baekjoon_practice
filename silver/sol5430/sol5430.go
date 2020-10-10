package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	br = bufio.NewReader(os.Stdin)
	bw = bufio.NewWriter(os.Stdout)
)

func main() {
	defer bw.Flush()
	var t, n int
	scanf("%d\n", &t)
	var p, xx string
	for i := 0; i < t; i++ {
		scanf("%s\n%d\n%s\n", &p, &n, &xx)
		x := make([]int, 0, n)
		for _, arg := range strings.Split(xx[1:len(xx)-1], ",") {
			number, _ := strconv.Atoi(arg)
			x = append(x, number)
		}

		h := true
		from, to := 0, n-1
		d := 0
		for _, v := range p {
			switch v {
			case 'R':
				h = !h
			case 'D':
				d++
				if h {
					from++
				} else {
					to--
				}
			}
		}
		if d == n {
			printf("[]\n")
			continue
		} else if d > n {
			printf("error\n")
			continue
		}
		if h {
			printf("[")
			for j := from; j < to; j++ {
				printf("%d,", x[j])
			}
			printf("%d", x[to])
			printf("]\n")
		} else {
			printf("[")
			for j := to; j > from; j-- {
				printf("%d,", x[j])
			}
			printf("%d", x[from])
			printf("]\n")
		}
	}
}
func printf(f string, a ...interface{}) { fmt.Fprintf(bw, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(br, f, a...) }
