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
	for i := 1; i <= n; i++ {
		for j := 1; j <= n-i; j++ {
			fmt.Fprint(bw, " ")
		}
		for j := 1; j < 2*i; j++ {
			fmt.Fprint(bw, "*")
		}
		fmt.Fprintln(bw)
	}
}
