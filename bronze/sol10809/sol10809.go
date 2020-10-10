package main

import (
	"fmt"
)

func main() {
	var (
		s  string
		al [26]int
	)
	fmt.Scan(&s)

	for i, c := range s {
		if al[c-'a'] == 0 {
			al[c-'a'] = i + 1
		}
	}
	for _, i := range al {
		fmt.Printf("%d ", i-1)
	}
}
