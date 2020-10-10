package main

import "fmt"

func main() {
	var t, k, n int
	fmt.Scan(&t)
	h := [15][15]int{}
	for i := 0; i < 15; i++ {
		for j := 1; j < 15; j++ {
			if i == 0 {
				h[i][j] = j
			} else {
				h[i][j] = h[i-1][j] + h[i][j-1]
			}
		}
	}
	for i := 0; i < t; i++ {
		fmt.Scan(&k, &n)
		fmt.Println(h[k][n])
	}
}
