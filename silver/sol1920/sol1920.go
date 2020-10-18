package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	sc = bufio.NewScanner(os.Stdin)
	a  map[int]bool
)

func main() {
	n := nextInt()
	a = make(map[int]bool)
	for i := 0; i < n; i++ {
		t := nextInt()
		a[t] = true
	}
	m := nextInt()
	for i := 0; i < m; i++ {
		mm := nextInt()
		if a[mm] {
			fmt.Println(1)
		} else {
			fmt.Println(0)
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
