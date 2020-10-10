package main

import (
	"bufio"
	"os"
	"strconv"
)

var br = bufio.NewReader(os.Stdin)
var sc = bufio.NewScanner(os.Stdin)
var bw = bufio.NewWriter(os.Stdout)

var (
	n, w  int
	load  = [1001][1001]int{}
	check = [1001][1001]int{}
	pol   [2]police
)

func main() {
	n, w = nextInt(), nextInt()
	for i := 0; i < w; i++ {
		a, b := nextInt(), nextInt()
		load[a][b] = 1
	}
	find()

}
func find() {

}
func nextInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}
func init() {
	sc.Split(bufio.ScanWords)
}

type police struct {
	x int
	y int
}
