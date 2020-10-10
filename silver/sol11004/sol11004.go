package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	n, k := nextInt(), nextInt()
	a := make([]int, n+1)
	for i := 0; i < n; i++ {
		a[i] = nextInt()
	}
	sort.Ints(a)
	fmt.Print(a[k-1])
}
func nextInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}
func init() {
	sc.Split(bufio.ScanWords)
}
