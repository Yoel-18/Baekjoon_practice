package main

import "fmt"

func main() {
	n, t := 0, [1000001]int{}
	fmt.Scan(&n)
	t[1] = 1
	t[2] = 2
	for i := 3; i <= n; i++ {
		t[i] = (t[i-1] + t[i-2]) % 15746
	}
	fmt.Print(t[n])
}
