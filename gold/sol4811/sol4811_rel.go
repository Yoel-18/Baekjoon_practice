package main

// import "fmt"

// var (
// 	n, count int
// 	dp       [31]int
// )

// func main() {
// 	dp[0] = 1
// 	dp[1] = 1
// 	dp[2] = 2
// 	for i := 3; i <= 30; i++ {
// 		for j := 0; j < i; j++ {
// 			dp[i] += dp[j] * dp[i-1-j]
// 		}
// 	}
// 	for {
// 		fmt.Scan(&n)
// 		if n == 0 {
// 			break
// 		}
// 		fmt.Println(dp[n])
// 	}
// }
