package main

import (
	"fmt"
	"strings"
)

func main() {
	var a, b string
	fmt.Scan(&a, &b)
	aa := strings.Split(a, "")
	bb := strings.Split(b, "")
	if bb[2] > aa[2] {
		fmt.Print(bb[2] + bb[1] + bb[0])
		return
	} else if bb[2] == aa[2] {
		if bb[1] > aa[1] {
			fmt.Print(bb[2] + bb[1] + bb[0])
			return
		} else if bb[1] == aa[1] {
			if bb[0] > aa[0] {
				fmt.Print(bb[2] + bb[1] + bb[0])
				return
			}
			fmt.Print(aa[2] + aa[1] + aa[0])
			return
		} else {
			fmt.Print(aa[2] + aa[1] + aa[0])
			return
		}
	} else {
		fmt.Print(aa[2] + aa[1] + aa[0])
		return
	}
}
