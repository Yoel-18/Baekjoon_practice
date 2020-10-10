package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	br = bufio.NewReader(os.Stdin)
	bw = bufio.NewWriter(os.Stdout)
	a  [46]int
	b  [46]int
	k  int
)

func main() {
	defer bw.Flush()
	fmt.Fscan(br, &k)
	a[0] = 1
	b[0] = 0
	for i := 1; i <= k; i++ {
		a[i] = b[i-1]
		b[i] = a[i-1] + b[i-1]
	}
	fmt.Fprint(bw, a[k], b[k])
}
