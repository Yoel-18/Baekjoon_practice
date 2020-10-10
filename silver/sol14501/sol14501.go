package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var br = bufio.NewReader(os.Stdin)
var sc = bufio.NewScanner(os.Stdin)
var bw = bufio.NewWriter(os.Stdout)

func main() {
	n := nextInt()
	t, p := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		t[i], p[i] = nextInt(), nextInt()
	}
	r := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		r[i] = r[i+1]
		if i+t[i] <= n {
			if a := r[i+t[i]] + p[i]; r[i] < a {
				r[i] = a
			}
		}
	}
	fmt.Println(r[0])
}
func nextInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}
func init() {
	sc.Split(bufio.ScanWords)
}
