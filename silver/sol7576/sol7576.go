package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// )

// var br = bufio.NewReader(os.Stdin)
// var sc = bufio.NewScanner(os.Stdin)
// var bw = bufio.NewWriter(os.Stdout)

// var (
// 	tomato      [1000][1000]int
// 	n, m, count int
// 	d           = [4]pos{{1, 0}, {-1, 0}, {0, -1}, {0, 1}}
// 	q           []status
// )

// func main() {
// 	defer bw.Flush()
// 	m, n = nextInt(), nextInt()
// 	q = make([]status, 0, m*n)
// 	for i := 0; i < n; i++ {
// 		for j := 0; j < m; j++ {
// 			tomato[i][j] = nextInt()
// 			if tomato[i][j] == 1 {
// 				q = append(q, status{pos{i, j}, 0})
// 			}
// 		}
// 	}
// 	find()
// 	isAll()
// }
// func find() {
// 	for len(q) > 0 {
// 		cur := q[0]
// 		q = q[1:]

// 		if cur.day > count {
// 			count = cur.day
// 		}

// 		for _, i := range d {
// 			new := pos{cur.x + i.x, cur.y + i.y}
// 			if new.x < 0 || new.x >= n || new.y < 0 || new.y >= m {
// 				continue
// 			}
// 			if tomato[new.x][new.y] != 0 {
// 				continue
// 			}
// 			tomato[new.x][new.y] = 1
// 			q = append(q, status{pos{new.x, new.y}, cur.day + 1})
// 		}
// 	}
// }
// func isAll() {
// 	for i := 0; i < n; i++ {
// 		for j := 0; j < m; j++ {
// 			if tomato[i][j] == 0 {
// 				fmt.Println(-1)
// 				return
// 			}
// 		}
// 	}
// 	fmt.Println(count)
// }

// type pos struct {
// 	x int
// 	y int
// }
// type status struct {
// 	pos
// 	day int
// }

// func nextInt() int {
// 	sc.Scan()
// 	n, _ := strconv.Atoi(sc.Text())
// 	return n
// }
// func init() {
// 	sc.Split(bufio.ScanWords)
// }
