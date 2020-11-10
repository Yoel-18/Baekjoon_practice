package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	k := 1000 - n
	count := 0
	coin := []int{500, 100, 50, 10, 5, 1}
	i := 0
	for k > 0 {
		if k >= coin[i] {
			count += k / coin[i]
			k %= coin[i]
		}
		i++
	}
	fmt.Print(count)
}
