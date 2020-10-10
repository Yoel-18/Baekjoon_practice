package main

import "fmt"

var (
	n, count int
	dp       [31][31]int
)

func main() {

	for {
		fmt.Scan(&n)
		if n == 0 {
			break
		}
		fmt.Println(pill(n-1, 1))
	}
}
func pill(w, h int) int {
	if w > 0 {
		if dp[w][h] != 0 {
			return dp[w][h]
		}
	}
	if w == 0 {
		dp[w][h] = 1
		return 1
	}
	count := 0
	count += pill(w-1, h+1)
	if h >= 1 {
		count += pill(w, h-1)
	}
	dp[w][h] = count
	return count
}
