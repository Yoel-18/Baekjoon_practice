package main

import (
	"bufio"
	"os"
	"strconv"
)

var (
	sc = bufio.NewScanner(os.Stdin)
	bw = bufio.NewWriter(os.Stdout)
)

//	정렬 안하고 map 이용...천재들 많네...
func main() {
	defer bw.Flush()
	n := nextInt()
	nn := make(map[int]bool)
	for i := 0; i < n; i++ {
		num := nextInt()
		nn[num] = true
	}
	m := nextInt()
	for i := 0; i < m; i++ {
		num := nextInt()
		if nn[num] {
			bw.WriteString("1 ")
		} else {
			bw.WriteString("0 ")
		}
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
