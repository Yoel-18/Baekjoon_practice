package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	sc = bufio.NewScanner(os.Stdin)
	n  int
	r  [100000]int
	o  [100000]int
)

func main() {
	n = nextInt()
	for i := 0; i < n-1; i++ {
		r[i] = nextInt()
	}
	for i := 0; i < n; i++ {
		o[i] = nextInt()
	}
	for i := 0; i < n-1; i++ {
		if o[i] < o[i+1] {
			o[i+1] = o[i]
		}
	}
	cost := 0
	for i := 0; i < n-1; i++ {
		cost += (o[i] * r[i])
	}
	fmt.Print(cost)
}
func nextInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}
func init() {
	sc.Split(bufio.ScanWords)
}
