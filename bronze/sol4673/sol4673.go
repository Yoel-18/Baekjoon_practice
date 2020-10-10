package main

import "fmt"

var (
	check [10001]bool
)

func main() {
	for i := 1; i <= 10000; i++ {
		a := d(i)
		if a <= 10000 {
			check[a] = true
		}

	}
	for i := 1; i <= 10000; i++ {
		if !check[i] {
			fmt.Println(i)
		}
	}
}
func d(n int) int {
	sum := n
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}
