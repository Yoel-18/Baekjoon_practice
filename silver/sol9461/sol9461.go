package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)
	var p [101]int
	for i := 0; i < t; i++ {
		var n int
		fmt.Scan(&n)
		p = [101]int{}
		p[1] = 1
		p[2] = 1
		p[3] = 1
		for j := 4; j <= n; j++ {
			p[j] = p[j-2] + p[j-3]
		}
		fmt.Println(p[n])
	}
}
