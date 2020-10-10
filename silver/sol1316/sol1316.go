package main

import "fmt"

func main() {
	var n int
	var str string
	fmt.Scan(&n)
	count := 0
	for i := 0; i < n; i++ {
		check := true
		fmt.Scan(&str)
		for j := 1; j < len(str); j++ {
			if str[j] != str[j-1] {
				for k := 0; k < j-1; k++ {
					if str[j] == str[k] {
						check = false
					}
				}
			}
		}
		if check {
			count++
		}
	}
	fmt.Print(count)
}
