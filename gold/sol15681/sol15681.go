package main

import (
	"bufio"
	"os"
	"strconv"
)

var (
	sc    = bufio.NewScanner(os.Stdin)
	bw    = bufio.NewWriter(os.Stdout)
	tree  [][]int
	check []bool
	dp    []int
)

func main() {
	defer bw.Flush()
	n, r, q := nextInt(), nextInt(), nextInt()
	tree = make([][]int, n+1)
	check = make([]bool, n+1)
	dp = make([]int, n+1)
	for i := 1; i < n; i++ {
		u, v := nextInt(), nextInt()
		tree[u] = append(tree[u], v)
		tree[v] = append(tree[v], u)
	}
	makeTree(r)
	for i := 0; i < q; i++ {
		uu := nextInt()
		str := strconv.Itoa(dp[uu])
		bw.WriteString(str + "\n")
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
