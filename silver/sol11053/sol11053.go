package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var br = bufio.NewReader(os.Stdin)
var sc = bufio.NewScanner(os.Stdin)
var bw = bufio.NewWriter(os.Stdout)

func main() {
	n := nextInt()
	a := make([]int, n)
	check := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = nextInt()
	}
	check[0] = 1
	for i := 1; i < n; i++ {
		check[i] = 1
		for j := 0; j < i; j++ {
			if a[j] < a[i] && check[i] <= check[j] {
				check[i] = check[j] + 1
			}
		}
	}
	max := 0
	for i := 0; i < n; i++ {
		if max < check[i] {
			max = check[i]
		}
	}
	fmt.Print(max)

}
func nextInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}
func init() {
	sc.Split(bufio.ScanWords)
}
