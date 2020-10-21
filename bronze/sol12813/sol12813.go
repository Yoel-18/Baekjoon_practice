package main

import (
	"bufio"
	"os"
)

var in = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func next() []byte {
	in.Scan()
	a := in.Bytes()
	r := make([]byte, len(a))
	copy(r, a)
	return r
}

func main() {
	in.Buffer(make([]byte, 131072), 131072)
	a, b := next(), next()
	for i := range a {
		if a[i] == '1' && b[i] == '1' {
			out.WriteByte('1')
		} else {
			out.WriteByte('0')
		}
	}
	out.WriteByte('\n')
	for i := range a {
		if a[i] == '1' || b[i] == '1' {
			out.WriteByte('1')
		} else {
			out.WriteByte('0')
		}
	}
	out.WriteByte('\n')
	for i := range a {
		if a[i] != b[i] {
			out.WriteByte('1')
		} else {
			out.WriteByte('0')
		}
	}
	out.WriteByte('\n')
	for i := range a {
		if a[i] == '0' {
			out.WriteByte('1')
		} else {
			out.WriteByte('0')
		}
	}
	out.WriteByte('\n')
	for i := range a {
		if b[i] == '0' {
			out.WriteByte('1')
		} else {
			out.WriteByte('0')
		}
	}
	out.Flush()
}
