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
	check            bool
)

//	동전의 갯수가 정해져있다고 가정
const coinCount = 10

func main() {
	defer bw.Flush()
	fmt.Fscan(br, &n, &k)
	for i := 0; i < n; i++ {
		fmt.Fscan(br, &a[i])
	}
	//	최솟값을 먼저 찾는다???
	for i := n - 1; i >= 0; i-- {
		if a[i] <= k {
			count += k / a[i]
			k %= a[i]
			cur = i
		}
	}
	//	최솟값과 같지 않으면
	for count != coinCount {
		k += a[cur]                     //	가장 마지막에 더한 값을 더해준다
		for i := cur - 1; i >= 0; i-- { //	더 작은 값으로 채워서 갯수를 즐려준다
			if a[i] <= k {
				count += k / a[i]
				k %= a[i]
				cur = i
			}
		}
		if count != coinCount { //	같지 않으면 다시 루프 반복
			continue
		}
		check = true //	같으면 체크를 해주고 반복문 종료????
	}

	if check {
		fmt.Fprint(bw, "YES", count)
	} else {
		fmt.Fprint(bw, "NO", count)
	}
}
