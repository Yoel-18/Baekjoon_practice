package main

import "fmt"

func main() {
	a := [5]int{}
	for i := 0; i < 5; i++ {
		fmt.Scan(&a[i])
		if a[i] < 40 {
			a[i] = 40
		}
	}
	fmt.Print((a[0] + a[1] + a[2] + a[3] + a[4]) / 5)
}
