package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	n := nextInt()
	num := [301]int{}
	dp := [301]int{}
	for i := 1; i <= n; i++ {
		num[i] = nextInt()
	}
	dp[1] = num[1]
	dp[2] = num[1] + num[2]
	for i := 3; i <= n; i++ {
		dp[i] = num[i] + maxInt(dp[i-3]+num[i-1], dp[i-2])
	}
	fmt.Print(dp[n])
}
func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func nextInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}
func init() {
	sc.Split(bufio.ScanWords)
}
