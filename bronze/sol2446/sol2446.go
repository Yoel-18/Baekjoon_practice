package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			fmt.Print(" ")
		}
		for k := 0; k < (2*n-1)-(2*i); k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	for i := 0; i < n-1; i++ {
		for j := 1; j < n-1-i; j++ {
			fmt.Print(" ")
		}
		for k := 0; k < 3+2*i; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
