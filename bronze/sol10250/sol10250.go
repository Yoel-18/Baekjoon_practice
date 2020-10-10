package main

import "fmt"

func main() {
	var t, h, w, n int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		fmt.Scan(&h, &w, &n)
		count := 0
		for n > h {
			n -= h
			count++
		}
		fmt.Println(n*100 + count + 1)
	}
}
