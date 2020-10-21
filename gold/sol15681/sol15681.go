package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	sc    = bufio.NewScanner(os.Stdin)
	tree  [][]int
	check = [10001]bool{}
	dp    = [10001]int{}
)

func main() {
	n, r, q := nextInt(), nextInt(), nextInt()
	tree = make([][]int, n+1)
	for i := 1; i < n; i++ {
		u, v := nextInt(), nextInt()
		tree[u] = append(tree[u], v)
		tree[v] = append(tree[v], u)
	}
	makeTree(r)
	for i := 0; i < q; i++ {
		uu := nextInt()
		fmt.Println(dp[uu])
	}
}
func makeTree(cur int) int {
	check[cur] = true
	for i := 0; i < len(tree[cur]); i++ {
		next := tree[cur][i]
		if !check[next] {
			dp[cur] += makeTree(next)
		}
	}
	dp[cur]++
	return dp[cur]
}
func nextInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}
func init() {
	sc.Split(bufio.ScanWords)
}
