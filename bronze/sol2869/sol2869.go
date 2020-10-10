package main

import "fmt"

func main() {
	var (
		a int
		b int
		v int
	)
	fmt.Scan(&a, &b, &v)
	count := (v - b) / (a - b)
	if (v-b)%(a-b) != 0 {
		count++
	}
	fmt.Print(count)
}
