package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	dp := [1001]int{}
	dp[1] = 1
	dp[2] = 3
	for i := 3; i <= n; i++ {
		dp[i] = (dp[i-1] + 2*dp[i-2]) % 10007
	}
	fmt.Print(dp[n])
}
