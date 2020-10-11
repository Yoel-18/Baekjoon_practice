package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	br          = bufio.NewReader(os.Stdin)
	bw          = bufio.NewWriter(os.Stdout)
	n, k, count int
	a           [10]int
)

func main() {
	defer bw.Flush()
	fmt.Fscan(br, &n, &k)
	for i := 0; i < n; i++ {
		fmt.Fscan(br, &a[i])
	}
	for i := n - 1; i >= 0; i-- {
		if a[i] <= k {
			count += k / a[i]
			k %= a[i]
		}
	}
	fmt.Fprint(bw, count)
}
