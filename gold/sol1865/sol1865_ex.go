package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var (
	sc          = bufio.NewScanner(os.Stdin)
	tc, n, m, w int
)

func main() {
	tc = nextInt()
	for i := 0; i < tc; i++ {
		n, m, w = nextInt(), nextInt(), nextInt()
		g := make([][]pos, m+1)
		for j := 0; j < m; j++ { //	도로 입력, 방향 x
			s, e, t := nextInt(), nextInt(), nextInt()
			g[s] = append(g[s], pos{e, t})
			g[e] = append(g[e], pos{s, t})
		}
		for j := 0; j < w; j++ { //	웜홀 입력, 방향 o
			s, e, t := nextInt(), nextInt(), nextInt()
			g[s] = append(g[s], pos{e, (-1) * t})
		}

		dis := make([]int, n+1)
		for j := range dis {
			dis[j] = math.MaxInt32
		}
		dis[1] = 0
		count := 0 //	들리는 횟수
		isUpdate := false

		//	count가 n보다 크면 음의 사이클 존재
		//	혹은 업데이트가 없으면 종료
		for isUpdate && count < n {
			isUpdate = false
			count++
			for i := 1; i <= n; i++ {
				for _, pos := range g[i] {
					if dis[i]+pos.cost < dis[pos.node] {
						dis[pos.node] = dis[i] + pos.cost
						isUpdate = true
					}
				}
			}
		}

		if count == n {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

type pos struct {
	node, cost int
}

func nextInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}
func init() {
	sc.Split(bufio.ScanWords)
}
