package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	check := [10]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	for {
		check[n%10]++
		if n/10 == 0 {
			break
		}
		n /= 10
	}

	num := 0
	for i := 0; i < 10; i++ {
		if i != 9 && i != 6 {
			num = maxInt(num, check[i])
		}
	}
	fmt.Print(maxInt(num, (check[6]+check[9]+1)/2))
}
func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
