package main

import "fmt"

func main() {
	var n, m int
	fmt.Scanf("%d %d", &n, &m)
	check := [1000001]bool{}
	check[1] = true
	for i := 2; i <= m; i++ {
		if check[i] {
			continue
		}
		for j := i + i; j <= m; j += i {
			check[j] = true
		}
	}
	for i := n; i <= m; i++ {
		if !check[i] {
			fmt.Println(i)
		}
	}
}
