package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	br := bufio.NewReader(os.Stdin)
	bw := bufio.NewWriter(os.Stdout)
	defer bw.Flush()
	var n int
	fmt.Fscan(br, &n)
	for j := n; j > 0; j-- {
		for k := n; k > j; k-- {
			fmt.Fprint(bw, " ")
		}
		for i := j; i > 0; i-- {
			fmt.Fprint(bw, "*")
		}
		fmt.Fprintln(bw)
	}
}
