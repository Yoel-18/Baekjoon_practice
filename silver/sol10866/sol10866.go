package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	br = bufio.NewReader(os.Stdin)
	bw = bufio.NewWriter(os.Stdout)
	n  int
)

//	리스트 사용...
func main() {
	defer bw.Flush()
	fmt.Fscan(br, &n)
	d := make([]int, 0)
	for i := 0; i < n; i++ {
		var str string
		var num int
		fmt.Fscan(br, &str)
		switch str {
		case "push_front":
			fmt.Fscan(br, &num)
		case "push_back":
			fmt.Fscan(br, &num)
			d = append(d, num)
		case "pop_front":
			if len(d) == 0 {
				fmt.Fprintln(bw, -1)
			} else {
				cur := d[0]
				d = d[1:]
				fmt.Fprintln(bw, cur)
			}
		case "pop_back":
			if len(d) == 0 {
				fmt.Fprintln(bw, -1)
			} else {
				cur := d[len(d)-1]
				d = d[:len(d)-1]
				fmt.Fprintln(bw, cur)
			}
		case "size":
			fmt.Fprintln(bw, len(d))
		case "empty":
			if len(d) != 0 {
				fmt.Fprintln(bw, 0)
			} else {
				fmt.Fprintln(bw, -1)
			}
		case "front":
			if len(d) == 0 {
				fmt.Fprintln(bw, -1)
			} else {
				fmt.Fprintln(bw, d[0])
			}
		case "back":
			if len(d) == 0 {
				fmt.Fprintln(bw, -1)
			} else {
				fmt.Fprintln(bw, d[len(d)])
			}
		}
	}
}
