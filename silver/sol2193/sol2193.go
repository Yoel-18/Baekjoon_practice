package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	dp := [91]int{}
	dp[1] = 1
	dp[2] = 1
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	fmt.Print(dp[n])
}
