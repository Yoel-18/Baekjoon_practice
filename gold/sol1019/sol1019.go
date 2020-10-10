package main

import "fmt"

func main() {
	var n, pre int
	fmt.Scan(&n)
	num := [10]int{}
	for i := 1; 0 < n; i *= 10 {
		now := n % 10
		n /= 10
		for j := 0; j < now; j++ {
			num[j] += (n + 1) * i
		}
		num[now] += n*i + pre + 1
		for j := now + 1; j < 10; j++ {
			num[j] += n * i
		}
		num[0] -= i
		pre += now * i
	}
	for i := 0; i < 10; i++ {
		fmt.Print(num[i], " ")
	}
}
