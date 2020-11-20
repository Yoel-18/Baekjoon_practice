package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	left, right, answer := 1, k, 0
	for left <= right {
		mid := (left + right) / 2
		count := getNum(mid, n)
		if count >= k {
			answer = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	fmt.Print(answer)
}
func getNum(x, n int) int {
	count := 0
	for i := 1; i <= n; i++ {
		count += minInt(x/i, n)
	}
	return count
}

func minInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}
