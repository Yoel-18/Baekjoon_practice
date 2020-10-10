package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	t := nextInt()
	fmt.Scan(&t)
	var dp [30][30]int
	for i := 0; i < t; i++ {
		n, m := nextInt(), nextInt()
		for j := 1; j <= m; j++ {
			dp[1][j] = j
		}
		for j := 2; j <= n; j++ {
			for k := 2; k <= m; k++ {
				dp[j][k] = dp[j][k-1] + dp[j-1][k-1]
			}
		}
		fmt.Println(dp[n][m])
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
