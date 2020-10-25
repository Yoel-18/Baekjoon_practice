package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if n == 1 || n == 3 {
		fmt.Print(-1)
	} else if n == 4 {
		fmt.Println(2)
	} else {
		count := 0
		count += n / 5
		n %= 5
		count += n / 2
		n %= 2
		if n == 1 {
			n += 5
			count += n/2 - 1
			n %= 2
		}
		fmt.Print(count)
	}
}
