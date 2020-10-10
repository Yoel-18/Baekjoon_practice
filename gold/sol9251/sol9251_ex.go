package main

import (
	"bufio"
	"fmt"
	"os"
)

var in = bufio.NewScanner(os.Stdin)

func next() []byte {
	in.Scan()
	return in.Bytes()
}

func sol9251() {
	p, q := next(), next()
	if len(p) == 0 || len(q) == 0 {
		fmt.Println("0")
		return
	}
	r, s := make([]int, len(q)), make([]int, len(q))
	for _, c := range p {
		for i, d := range q {
			if i == 0 {
				if c == d {
					s[0] = 1
				}
				if s[0] < r[0] {
					s[0] = r[0]
				}
				continue
			}
			s[i] = r[i-1]
			if c == d {
				s[i]++
			}
			if s[i] < r[i] {
				s[i] = r[i]
			}
			if s[i] < s[i-1] {
				s[i] = s[i-1]
			}
		}
		r, s = s, r
	}
	fmt.Println(r[len(r)-1])
}
