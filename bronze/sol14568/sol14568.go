package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	count := 0
	for i := 1; i <= 100; i++ {
		for j := 1; j <= 100; j++ {
			for k := 1; k <= 100; k++ {
				if i+j+k == n && i >= j+2 && k%2 == 0 {
					count++
				}
			}
		}
	}
	fmt.Print(count)
}
