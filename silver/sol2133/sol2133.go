package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	dp := [31]int{}
	if n > 1 && n%2 == 0 {
		dp[0] = 1
		dp[2] = 3
		for i := 4; i <= n; i += 2 {
			dp[i] = dp[2] * dp[i-2]
			for j := 4; j <= i; j += 2 {
				dp[i] += 2 * dp[i-j]
			}
		}
	}
	fmt.Print(dp[n])
}
