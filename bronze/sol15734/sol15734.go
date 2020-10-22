package main

import "fmt"

func main() {
	var l, r, a int
	fmt.Scan(&l, &r, &a)
	answer := 0
	for i := 0; i <= a; i++ {
		answer = maxInt(answer, minInt(l+i, r+a-i))
	}
	fmt.Print(answer * 2)
}
func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func minInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}
