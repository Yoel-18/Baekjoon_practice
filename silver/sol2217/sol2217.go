package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	br := bufio.NewReader(os.Stdin)
	bw := bufio.NewWriter(os.Stdout)
	defer bw.Flush()
	var n int
	fmt.Fscan(br, &n)
	l := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(br, &l[i])
	}
	sort.Ints(l)
	max := 0
	for i := n - 1; i >= 0; i-- {
		l[i] = l[i] * (n - i)
		if max < l[i] {
			max = l[i]
		}
	}
	fmt.Fprint(bw, max)
}
