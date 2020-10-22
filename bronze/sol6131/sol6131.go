package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	count := 0
	for i := 1; i <= 1000; i++ {
		for j := 1; j <= 1000; j++ {
			if i*i == j*j+n {
				count++
			}
		}
	}
	fmt.Print(count)
}
