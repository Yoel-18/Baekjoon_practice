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

//	많이 느림
func main() {
	defer bw.Flush()
	var n, x int
	fmt.Fscan(br, &n, &x)
	if n > x || x > n*26 {
		fmt.Fprint(bw, "!")
		return
	}
	answer := make([]byte, 0)
	for i := 0; i < n; i++ {
		answer = append(answer, 'A')
	}
	x -= n

	for i := n - 1; i >= 0 && x > 0; i-- {
		cur := minInt(x, 25)
		answer[i] += byte(cur)
		x -= cur
	}
	for _, s := range answer {
		fmt.Fprint(bw, string(s))
	}
}
func minInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}
