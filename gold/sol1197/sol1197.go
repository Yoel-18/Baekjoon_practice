package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var (
	sc     = bufio.NewScanner(os.Stdin)
	v, e   int
	g      []tree
	parent []int
)

//	최소 신장 트리
//	크루스칼 알고리즘 이용

func main() {
	v, e = nextInt(), nextInt()
	parent = make([]int, v+1)
	for i := 1; i <= v; i++ {
		parent[i] = i
	}
	for i := 0; i < e; i++ {
		a, b, c := nextInt(), nextInt(), nextInt()
		g = append(g, tree{a, b, c})
	}
	//	그래프를 비용이 적은 순서로 정렬
	sort.Slice(g, func(i, j int) bool {
		return g[i].cost < g[j].cost
	})

	sum, edgeCount := 0, 0
	//	최소 비용의 간선부터 시작
	for i := 0; i < e; i++ {
		from := g[i].from
		to := g[i].to
		cost := g[i].cost

		if find(from) != find(to) { //	두 정점이 연결되어 있지 않다면
			union(from, to) //	두 정점을 연결
			sum += cost
			edgeCount++
			if edgeCount == v-1 {
				break
			}
		}
	}
	fmt.Print(sum)
}

//	해당 정점을 중심으로 연결된 정점 찾기?
func find(x int) int {
	if x == parent[x] {
		return x
	}
	root := find(parent[x])
	parent[x] = root
	return root
}

//	두 정점 연결
func union(x, y int) {
	x = find(x)
	y = find(y)

	parent[y] = x
}

type tree struct {
	from, to, cost int
}

func nextInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}
func init() {
	sc.Split(bufio.ScanWords)
}
