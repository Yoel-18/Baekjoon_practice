package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	sc   = bufio.NewScanner(os.Stdin)
	n, k int
	a    = [101]int{}
	dp   = [10001]int{}
)

func main() {
	n, k = nextInt(), nextInt()
	for i := 1; i <= n; i++ {
		a[i] = nextInt()
	}
	dp[0] = 1
	for i := 1; i <= n; i++ {
		for j := a[i]; j <= k; j++ {
			dp[j] += dp[j-a[i]]
		}
	}
	fmt.Print(dp[k])
}
func nextInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}
func init() {
	sc.Split(bufio.ScanWords)
}

type coin struct {
	cnt       int
	coinCount int
}
