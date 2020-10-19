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
	var n, b, num int
	var str string
	fmt.Fscan(br, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(br, &str)
		switch str {
		case "add":
			fmt.Fscan(br, &num)
			b = b | (1 << num)
		case "remove":
			fmt.Fscan(br, &num)
			b = b & ^(1 << num)
		case "check":
			fmt.Fscan(br, &num)
			if b&(1<<num) != 0 {
				fmt.Fprintln(bw, 1)
			} else {
				fmt.Fprintln(bw, 0)
			}
		case "toggle":
			fmt.Fscan(br, &num)
			b = b ^ (1 << num)
		case "all":
			b = ^0
		case "empty":
			b = 0
		}
	}
}
