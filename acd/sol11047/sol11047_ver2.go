package main

import (
	"bufio"
	"fmt"
	"os"
)

//	동전의 갯수가 정해져있다고 가정
const coinCount = 10

var (
	br               = bufio.NewReader(os.Stdin)
	bw               = bufio.NewWriter(os.Stdout)
	n, k, count, cur int
	a                [10]int
)

func main() {
	defer bw.Flush()
	fmt.Fscan(br, &n, &k)
	for i := 0; i < n; i++ {
		fmt.Fscan(br, &a[i])
	}
	for i := 1; i <= coinCount; i++ {
		for j := 1; j <= coinCount; j++ {

		}
	}

	if count == coinCount {
		fmt.Fprint(bw, "YES", count)
	} else {
		fmt.Fprint(bw, "NO", count)
	}
}
