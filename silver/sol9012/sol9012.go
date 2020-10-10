package main

import (
	"fmt"
	"strings"
)

var (
	t     int
	str   string
	vps   []string
	check bool
)

func main() {
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		fmt.Scan(&str)
		vps = strings.Split(str, "")
		check = true
		st := make([]string, len(vps))
		ptr := 0
		for j := 0; j < len(vps); j++ {
			if vps[j] == "(" {
				st = append(st, vps[j])
				ptr++
			} else {
				if ptr == 0 {
					check = false
					break
				}
				ptr--
			}
		}
		fmt.Println(st)
		if ptr == 0 && check {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
