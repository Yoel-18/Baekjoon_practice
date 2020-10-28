package main

import (
	"bufio"
	"os"
	"strconv"
)

var (
	sc   = bufio.NewScanner(os.Stdin)
	n, m int
	maap [10][10]int
)

//	다리 길이 2 이상
//	방향 바뀌면 안됨
func main() {
	n, m = nextInt(), nextInt()
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			maap[i][j] = nextInt()
		}
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
