package main

import "fmt"

func main() {
	fmt.Print(solution(5))
}

var (
	fib [100001]int
)

func solution(n int) int {
	fib[1] = 1
	for i := 2; i <= n; i++ {
		fib[i] = (fib[i-1] + fib[i-2]) % 1234567
	}
	answer := fib[n]
	return answer
}
