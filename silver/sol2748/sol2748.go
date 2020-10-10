package main

import "fmt"

var fib = [91]int{}

func main() {
	var n int
	fmt.Scan(&n)
	fib[0] = 0
	fib[1] = 1
	for i := 2; i <= n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	fmt.Print(fib[n])
}
