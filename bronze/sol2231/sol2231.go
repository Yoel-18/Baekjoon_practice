package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)
	s := len(strconv.Itoa(n))
	start := n - (9 * s)
	answer := 0

	for i := start; i < n; i++ {
		sum := i
		k := i
		for k > 0 {
			sum += k % 10
			k /= 10
		}
		if sum == n {
			answer = i
			break
		}
	}
	fmt.Print(answer)
}
