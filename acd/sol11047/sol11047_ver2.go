package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	br               = bufio.NewReader(os.Stdin)
	bw               = bufio.NewWriter(os.Stdout)
	n, k, count, cur int
	a                [10]int
)

//	동전의 갯수가 정해져있다고 가정
const coinCount = 10

func main() {
	defer bw.Flush()
	fmt.Fscan(br, &n, &k)
	for i := 0; i < n; i++ {
		fmt.Fscan(br, &a[i])
	}

	if count == coinCount {
		fmt.Fprint(bw, "YES", count)
	} else {
		fmt.Fprint(bw, "NO", count)
	}
}
