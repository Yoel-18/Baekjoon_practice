package main

import "fmt"

func main() {
	fib := [46]int{}
	var n int
	fmt.Scan(&n)
	fib[1] = 1
	fib[2] = 1
	for i := 3; i <= n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	fmt.Print(fib[n])
}
