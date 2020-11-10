package main

import (
	"bufio"
	"fmt"
	"strconv"

	"os"
)

var (
	sc = bufio.NewScanner(os.Stdin)
)

func main() {
	t := nextInt()
	for i := 0; i < t; i++ {
		n := nextInt()
		nn := make([]int, n+1)
		for j := 0; j < n; j++ {
			a, b := nextInt(), nextInt()
			nn[a] = b
		}
		answer := 1
		standrad := nn[1]
		for k := 2; k <= n; k++ {
			if standrad > nn[k] {
				answer++
				standrad = nn[k]
			}
		}
		fmt.Println(answer)
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
