package main

import "fmt"

func main() {
	var a, b, n, w int
	fmt.Scan(&a, &b, &n, &w)
	count := 0
	aa, bb := 0, 0
	for i := 1; i < n; i++ {
		for j := 1; j < n; j++ {
			if (i*a+j*b) == w && i+j == n {
				count++
				aa, bb = i, j
			}
		}
	}
	if count == 1 {
		fmt.Print(aa, bb)
	} else {
		fmt.Print(-1)
	}
}
