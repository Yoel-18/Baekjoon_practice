package main

import "fmt"

func sol1929_c() {
	var m, n int
	fmt.Scanf("%d %d", &m, &n)
	p := [1000001]bool{}
	p[1] = true
	for i := 2; i*i <= n; i++ {
		if !p[i] {
			for j := i * 2; j <= n; j += i {
				p[j] = true
			}
		}
	}
	for i := m; i <= n; i++ {
		if !p[i] {
			fmt.Println(i)
		}
	}
}
