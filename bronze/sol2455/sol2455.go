package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	br = bufio.NewReader(os.Stdin)
	bw = bufio.NewWriter(os.Stdout)
)

func main() {
	defer bw.Flush()
	var a, b int
	max := 0
	fmt.Fscanln(br, &a, &max)
	for i := 0; i < 3; i++ {
		fmt.Fscanln(br, &a, &b)
		cur := b - a
		if cur > 0 {
			max += cur
		}
	}
	fmt.Fprint(bw, max)
}
