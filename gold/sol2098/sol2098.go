package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var (
	sc  = bufio.NewScanner(os.Stdin)
	n   int
	w   [16][16]int
	dp  [16][1 << 16]int
	max = math.MaxInt64
)

func main() {
	n = nextInt()
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			w[i][j] = nextInt()
		}
	}
	fmt.Println(solve(0, 1))
}
func solve(node, visit int) int {
	if visit == (1<<n)-1 {
		if w[node][0] != 0 {
			return w[node][0]
		}
		return max
	}
	if dp[node][visit] != 0 {
		return dp[node][visit]
	}
	answer := max
	for i := 0; i < n; i++ {
		next := visit | (1 << i)
		if (visit&(1<<i)) == 0 && w[node][i] != 0 {
			answer = minInt(answer, w[node][i]+solve(i, next))
		}
	}
	dp[node][visit] = answer
	return dp[node][visit]
}
func nextInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}
func init() {
	sc.Split(bufio.ScanWords)
}
func minInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}
