package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	var r [26]int
	b, _ := ioutil.ReadAll(os.Stdin)
	for _, c := range b {
		switch {
		case 'A' <= c && c <= 'Z':
			r[c-'A']++
		case 'a' <= c && c <= 'z':
			r[c-'a']++
		}
	}
	q, m := int('?'), 0
	for c, v := range r[:] {
		if m < v {
			q, m = c+'A', v
		} else if m == v {
			q = '?'
		}
	}
	fmt.Printf("%c\n", q)
}
