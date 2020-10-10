package main

import (
	"bufio"
	"os"
	"strconv"
)

var br = bufio.NewReader(os.Stdin)
var sc = bufio.NewScanner(os.Stdin)
var bw = bufio.NewWriter(os.Stdout)

func main() {

}
func nextInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}
func init() {
	sc.Split(bufio.ScanWords)
}
