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
	n := 0
	fmt.Fscan(br, &n)
	num := make([]int, n)
	for i := range num {
		fmt.Fscan(br, &num[i])
	}
	sort.Ints(num)
	for _, number := range num {
		fmt.Fprintln(bw, number)
	}
}
