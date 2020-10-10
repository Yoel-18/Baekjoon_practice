package main

import "fmt"

var (
	f      = [1500000]int{}
	div    = 1000000
	pisano = 1500000
)

func main() {
	var n int
	fmt.Scan(&n)
	n %= pisano
	f[1] = 1
	for i := 2; i < pisano; i++ {
		f[i] = (f[i-2] + f[i-1]) % div
	}
	fmt.Print(f[n])
}
