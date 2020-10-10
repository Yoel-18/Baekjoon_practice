package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	a := make([][]int, 3, 3)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			tmp := nextInt()
			a[i] = append(a[i], tmp)
		}
	}
	fmt.Println(a)
}

var reader = bufio.NewReader(os.Stdin)
var scanner = bufio.NewScanner(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

type pos struct {
	i, j int
}

type status struct {
	pos
	day int
}

func init() {
	scanner.Split(bufio.ScanWords)
}

func nextInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
