package main

import "fmt"

func main() {
	var str string
	var num int
	fmt.Scan(&str)
	for i := 0; i < len(str); i++ {
		num *= 16
		if str[i] >= 'A' {
			num += (int)(str[i] - 'A' + 10)
		} else {
			num += (int)(str[i] - '0')
		}
	}
	fmt.Print(num)
}
