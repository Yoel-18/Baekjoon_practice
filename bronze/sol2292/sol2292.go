package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := 1
	count := 1
	for i := 1; a < n; i++ {
		a += 6 * i
		count++
	}
	fmt.Print(count)
}
