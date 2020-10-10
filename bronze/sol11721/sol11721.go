package main

import "fmt"

func main() {
	var n string
	fmt.Scan(&n)
	for 10 <= len(n) {
		fmt.Println(n[:10])
		n = n[10:]
	}
	fmt.Println(n)
}
