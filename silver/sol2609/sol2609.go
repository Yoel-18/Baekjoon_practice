package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	c, d := 0, 0
	if a > b {
		for i := a; ; i++ {
			if i%a == 0 && i%b == 0 {
				c = i
				break
			}
		}
		for j := b; ; j-- {
			if a%j == 0 && b%j == 0 {
				d = j
				break
			}
		}
	} else {
		for i := b; ; i++ {
			if i%b == 0 && i%a == 0 {
				c = i
				break
			}
		}
		for j := a; ; j-- {
			if a%j == 0 && b%j == 0 {
				d = j
				break
			}
		}
	}
	fmt.Println(d)
	fmt.Print(c)
}

// g, (a*b)%g
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
