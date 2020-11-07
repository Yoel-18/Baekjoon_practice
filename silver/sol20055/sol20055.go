package main

import (
	"bufio"
	"os"
	"strconv"
)

var (
	sc   = bufio.NewScanner(os.Stdin)
	n, k int
	a    [1001]int
)

func main() {
	n, k = nextInt(), nextInt()
	for i := 1; i <= 2*n; i++ {
		a[i] = nextInt()
	}
}
func nextInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}
func init() {
	sc.Split(bufio.ScanWords)
}
