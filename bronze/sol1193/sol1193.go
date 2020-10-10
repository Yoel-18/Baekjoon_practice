package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	count1, count2 := 1, 0
	for {
		if x <= count1+count2 {
			if count1%2 == 1 {
				fmt.Print((count1 - (x - count2 - 1)), "/", (x - count2))
				break
			} else {
				fmt.Print((x - count2), "/", (count1 - (x - count2 - 1)))
				break
			}
		} else {
			count2 += count1
			count1++
		}
	}
}
