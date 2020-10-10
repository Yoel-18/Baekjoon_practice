package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var (
	sc = bufio.NewScanner(os.Stdin)
	n  int
	p  []int
)

func main() {
	n = nextInt()
	p = make([]int, n+1)
	for i := 1; i <= n; i++ {
		p[i] = nextInt()
	}
	sort.Ints(p)
	sum := 0
	for i := 1; i <= n; i++ {
		sum += (p[i] * (n - i + 1))
	}
	fmt.Print(sum)
}

func nextInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}
func init() {
	sc.Split(bufio.ScanWords)
}
