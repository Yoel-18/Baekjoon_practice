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
	var n int
	fmt.Fscan(br, &n)
	for i := 1; i < n; i++ {
		for j := n; j > i; j-- {
			fmt.Fprint(bw, " ")
		}
		for j := 1; j < 2*i; j++ {
			fmt.Fprint(bw, "*")
		}
		fmt.Fprintln(bw)
	}
	for i := 1; i < 2*n; i++ {
		fmt.Fprint(bw, "*")
	}
	fmt.Fprintln(bw)
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			fmt.Fprint(bw, " ")
		}
		for j := 1; j < 2*(n-i); j++ {
			fmt.Fprint(bw, "*")
		}
		fmt.Fprintln(bw)
	}
}
