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
	var n, max int
	fmt.Scan(&n)
	a, dp := [1000]int{}, [1000]int{}
	for i := 0; i < n; i++ {
		a[i] = nextInt()
	}
	dp[0] = a[0]
	for i := 1; i < n; i++ {
		dp[i] = a[i]
		for j := 0; j < i; j++ {
			if a[i] > a[j] && dp[i] < dp[j]+a[i] {
				dp[i] = dp[j] + a[i]
			}
		}
	}
	for i := 0; i < n; i++ {
		if max < dp[i] {
			max = dp[i]
		}
	}
	fmt.Print(max)
}
func nextInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}
func init() {
	sc.Split(bufio.ScanWords)
}
