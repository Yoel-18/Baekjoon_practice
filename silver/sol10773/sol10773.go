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
	k := nextInt()
	st := make([]int, 0)
	for i := 0; i < k; i++ {
		num := nextInt()
		if num == 0 {
			st = st[:len(st)-1]
		} else {
			st = append(st, num)
		}
	}
	sum := 0
	for _, number := range st {
		sum += number
	}
	fmt.Print(sum)
}
func nextInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}
func init() {
	sc.Split(bufio.ScanWords)
}
