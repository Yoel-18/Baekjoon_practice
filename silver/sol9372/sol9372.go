package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	sc = bufio.NewScanner(os.Stdin)
)

func main() {
	t := nextInt()
	for i := 0; i < t; i++ {
		n, m := nextInt(), nextInt()
		for j := 0; j < m; j++ {
			a, b := nextInt(), nextInt()
			a = b
			b = a
		}
		fmt.Println(n - 1)
	}
}
func nextInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}
func init() {
	sc.Split(bufio.ScanWords)
}
