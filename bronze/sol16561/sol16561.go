package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	count := 0
	for i := 3; i <= 3000-6; i += 3 {
		for j := 3; j <= 3000-6; j += 3 {
			for k := 3; k <= 3000-6; k += 3 {
				if i+j+k == n {
					count++
				}
			}
		}
	}
	// n = n/3 - 3
	// if n < 0 {
	// 	count = 0
	// } else {
	// 	count = (n + 2) * (n + 1) / 2
	// }
	fmt.Print(count)
}
