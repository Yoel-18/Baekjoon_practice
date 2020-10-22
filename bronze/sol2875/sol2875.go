package main

import "fmt"

func main() {
	var n, m, k int
	fmt.Scan(&n, &m, &k)
	max := n / 2
	if n/2 > m {
		max = m
	}
	k -= n + m - max*3
	for k > 0 {
		k -= 3
		max--
	}
	fmt.Print(max)
}
