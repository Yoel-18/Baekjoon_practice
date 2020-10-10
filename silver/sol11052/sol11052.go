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
	p := make([]int, n+1)
	dp := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		p[i] = nextInt()
	}
	for i := 1; i < n+1; i++ {
		for j := 1; j <= i; j++ {
			if dp[i] < dp[i-j]+p[j] {
				dp[i] = dp[i-j] + p[j]
			}
		}
	}
	fmt.Print(dp[n])
}
func nextInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}
func init() {
	sc.Split(bufio.ScanWords)
}
