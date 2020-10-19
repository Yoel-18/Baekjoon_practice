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
	arr [500][2]int
	dp  [500][500]int
	max = math.MaxInt64
)

func main() {
	n = nextInt()
	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = -1
		}
	}
	for i := 0; i < n; i++ {
		arr[i][0], arr[i][1] = nextInt(), nextInt()
	}
	sol(0, n-1)
	fmt.Print(dp[0][n-1])

}
func sol(x, y int) int {
	if x == y {
		return 0
	}
	if dp[x][y] != -1 {
		return dp[x][y]
	}
	minNum := max
	for i := x; i < y; i++ {
		minNum = min(minNum, sol(x, i)+sol(i+1, y)+arr[x][0]*arr[i][1]*arr[y][1])
	}
	dp[x][y] = minNum
	return dp[x][y]
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func nextInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}
func init() {
	sc.Split(bufio.ScanWords)
}
