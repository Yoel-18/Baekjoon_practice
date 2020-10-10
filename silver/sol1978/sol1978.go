package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	sc  = bufio.NewScanner(os.Stdin)
	num [100]int
	p   [1001]bool
)

func main() {
	n := nextInt()
	p[1] = true
	count := 0
	for i := 2; i <= 1000; i++ {
		if p[i] {
			continue
		}
		for j := i + i; j <= 1000; j += i {
			p[j] = true
		}
	}
	for i := 0; i < n; i++ {
		num[i] = nextInt()
	}
	for i := 0; i < n; i++ {
		if !p[num[i]] {
			count++
		}
	}
	fmt.Print(count)

}
func nextInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}
func init() {
	sc.Split(bufio.ScanWords)
}
