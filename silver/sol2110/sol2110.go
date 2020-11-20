package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	br = bufio.NewReader(os.Stdin)
	bw = bufio.NewWriter(os.Stdout)
)

func main() {
	defer bw.Flush()
	var n, c int
	fmt.Fscan(br, &n, &c)
	x := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(br, &x[i])
	}
	sort.Ints(x) // 거리순 정렬
	left, right, dis, answer := 1, x[n-1]-x[0], 0, 0

	for left <= right { // 최대 거리를 타겟
		mid := (left + right) / 2
		cur, count := x[0], 1

		for i := 1; i < n; i++ {
			dis = x[i] - cur
			//	타겟보다 최대거리가 크면 설치
			if mid <= dis {
				count++
				cur = x[i]
			}
		}
		if count >= c {
			answer = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	fmt.Fprint(bw, answer)
}
