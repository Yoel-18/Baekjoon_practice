package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var (
	sc   = bufio.NewScanner(os.Stdin)
	n, k int
	dp   []int
	c    = [101]int{}
	max  = math.MaxInt32
)

func main() {
	n, k = nextInt(), nextInt()
	for i := 1; i <= n; i++ {
		c[i] = nextInt()
	}
	dp = make([]int, k+1)
	for i := 1; i <= k; i++ {
		dp[i] = max
	}
	for i := 1; i <= n; i++ {
		for j := c[i]; j <= k; j++ {
			if dp[j] > dp[j-c[i]]+1 {
				dp[j] = dp[j-c[i]] + 1
			}
		}
	}
	if dp[k] == max {
		fmt.Print(-1)
	} else {
		fmt.Print(dp[k])
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
