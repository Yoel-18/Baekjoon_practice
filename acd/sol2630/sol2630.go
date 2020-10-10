package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)
var (
	paper [][]int
	white int
	blue  int
)

func scanInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}

//Sol2630 function
func main() {
	sc.Split(bufio.ScanWords)
	n := scanInt()
	paper = make([][]int, n)

	for i := 0; i < n; i++ {
		paper[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&paper[i][j])
		}
	}

	find(0, 0, n)
	wr.WriteString(strconv.Itoa(white))
	wr.WriteByte('\n')
	wr.WriteString(strconv.Itoa(blue))
	wr.WriteByte('\n')
	wr.Flush()
}

func find(a, b, size int) {
	count := 0
	for i := a; i < a+size; i++ {
		for j := b; j < b+size; j++ {
			if paper[i][j] != 1 {
				count++
			}
		}
	}

	if count == size*size {
		white++
	} else if count == 0 {
		blue++
	} else {
		find(a, b, size/2)
		find(a+size/2, b, size/2)
		find(a, b+size/2, size/2)
		find(a+size/2, b+size/2, size/2)
	}
}
