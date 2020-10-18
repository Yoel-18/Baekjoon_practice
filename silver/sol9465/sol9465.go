package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	sc = bufio.NewScanner(os.Stdin)
	s  [2][100000]int
)

func main() {
	t := nextInt()
	for i := 0; i < t; i++ {
		n := nextInt()
		dp := [2][100000]int{}
		for j := 0; j < 2; j++ {
			for k := 0; k < n; k++ {
				s[j][k] = nextInt()
			}
		}
		dp[0][0] = s[0][0]
		dp[1][0] = s[1][0]
		for j := 2; j < n; j++ {
			dp[0][j] = maxInt(dp[1][j-1], dp[1][j-2]) + s[0][j]
			dp[1][j] = maxInt(dp[0][j-1], dp[0][j-2]) + s[1][j]
		}
		fmt.Println(maxInt(dp[0][n-1], dp[1][n-1]))
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
func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
