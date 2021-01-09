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
	for j := 1; j < n; j++ {
		for k := 0; k < j; k++ {
			fmt.Fprint(bw, "*")
		}
		fmt.Fprintln(bw)
	}
	for j := 1; j <= n; j++ {
		fmt.Fprint(bw, "*")
	}
	fmt.Fprintln(bw)
	for j := n; j > 1; j-- {
		for k := j; k > 1; k-- {
			fmt.Fprint(bw, "*")
		}
		fmt.Fprintln(bw)
	}
}
