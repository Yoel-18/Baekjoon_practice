package main

import "fmt"

func main() {
	var s, j, h, coke, cider int
	var a, b int
	fmt.Scan(&s)
	a = s
	fmt.Scan(&j)
	if a > j {
		a = j
	}
	fmt.Scan(&h)
	if a > h {
		a = h
	}
	fmt.Scan(&coke)
	b = coke
	fmt.Scan(&cider)
	if b > cider {
		b = cider
	}
	fmt.Print(a + b - 50)
}
