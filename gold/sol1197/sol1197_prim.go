package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// )

// var (
// 	sc    = bufio.NewScanner(os.Stdin)
// 	v, e  int
// 	g     []tree
// 	check []int
// )

// //	최소 신장 트리
// //	크루스칼 알고리즘 이용

// func main() {
// 	v, e = nextInt(), nextInt()
// 	parent = make([]int, v+1)
// 	for i := 1; i <= v; i++ {
// 		parent[i] = i
// 	}
// 	for i := 0; i < e; i++ {
// 		a, b, c := nextInt(), nextInt(), nextInt()
// 		g = append(g, tree{a, b, c})
// 	}

// 	sum, edgeCount := 0, 0

// 	fmt.Print(sum)
// }
// func findPrim() {
// 	ans := 0

// }

// type tree struct {
// 	from, to, cost int
// }

// func nextInt() int {
// 	sc.Scan()
// 	n, _ := strconv.Atoi(sc.Text())
// 	return n
// }
// func init() {
// 	sc.Split(bufio.ScanWords)
// }
