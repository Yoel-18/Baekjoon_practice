package main

import (
	"bufio"
	"container/heap"
	"os"
	"strconv"
)

var (
	sc *bufio.Scanner
	wr *bufio.Writer
)

func init() {
	sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	wr = bufio.NewWriter(os.Stdout)
}

func scanInt() int {
	sc.Scan()
	ret, _ := strconv.Atoi(sc.Text())
	return ret
}

func main() {
	defer wr.Flush()
	n, m := scanInt(), scanInt()
	s := make([][]bool, n)
	for i := range s {
		s[i] = make([]bool, m)
	}
	v := make([][]int, n)
	for i := range v {
		v[i] = make([]int, m)
		for j := range v[i] {
			v[i][j] = big
		}
	}

	for i := 0; i < n; i++ {
		sc.Scan()
		ss := sc.Text()
		for j := range ss {
			if ss[j] == '0' {
				s[i][j] = true
			} else {
				s[i][j] = false
			}
		}
	}
	h := &pp{{0, 0, 1, 0}}
	heap.Init(h)
	min := big
	for h.Len() > 0 {
		o := heap.Pop(h).(pair)
		if o.x == n-1 && o.y == m-1 {
			if o.z < min {
				min = o.z
			}
		}
		for i := 0; i < 4; i++ {
			xx, yy := o.x+dx[i], o.y+dy[i]
			if in(xx, n) && in(yy, m) {
				if s[xx][yy] && v[xx][yy] > o.z+1 {
					v[xx][yy] = o.z + 1
					heap.Push(h, pair{xx, yy, o.z + 1, o.w})
				} else if !s[xx][yy] && o.w == 0 && v[xx][yy] > o.z+1 {
					v[xx][yy] = o.z + 1
					heap.Push(h, pair{xx, yy, o.z + 1, 1})
				}
			}
		}
	}
	if min == big {
		wr.WriteString("-1")
	} else {
		wr.WriteString(strconv.Itoa(min))
	}
}

func in(x, n int) bool {
	return 0 <= x && x < n
}

const big = int(1<<31 - 1)

var (
	dx = []int{-1, 1, 0, 0}
	dy = []int{0, 0, -1, 1}
)

type (
	pair struct{ x, y, z, w int }
	pp   []pair
)

func (p pp) Len() int            { return len(p) }
func (p pp) Swap(i, j int)       { p[i], p[j] = p[j], p[i] }
func (p pp) Less(i, j int) bool  { return p[i].w < p[j].w || (p[i].w == p[j].w && p[i].z < p[j].z) }
func (p *pp) Push(x interface{}) { *p = append(*p, x.(pair)) }
func (p *pp) Pop() interface{} {
	old := *p
	n := len(old)
	x := old[n-1]
	*p = old[:n-1]
	return x
}

// package main
// import (
//     "bufio"
//     "os"
//     "strconv"
// )

// var sc = bufio.NewScanner(os.Stdin)
// var w = bufio.NewWriter(os.Stdout)
// func scanInt() int {
//     sc.Scan()
//     n,_ := strconv.Atoi(sc.Text())
//     return n
// }
// func scanText() string {
//     sc.Scan()
//     return sc.Text()
// }

// type Status struct {
//     i int
//     j int
//     breakWall bool
//     step int
// }

// var maze [][]int
// var visit [][][]bool
// var queue []Status
// var n,m int

// func main() {
//     sc.Split(bufio.ScanWords)
//     n = scanInt()
//     m = scanInt()
//     maze = make([][]int, n)
//     visit = make([][][]bool, n)
//     queue = make([]Status, 0)
//     for i:=0;i<n;i++ {
//         maze[i] = make([]int, m)
//         visit[i] = make([][]bool, m)
//         line := scanText()
//         for j,c := range line {
//             maze[i][j] = int(c-'0')
//         }
//         for k:=0;k<m;k++ {
//             visit[i][k] = make([]bool, 2)
//         }
//     }

//     w.WriteString(strconv.Itoa(bfs()))
//     w.WriteByte('\n')
//     w.Flush()
// }

// func bfs() int {
//     queue = append(queue,Status{i:0, j:0})
//     visit[0][0][0] = true
//     for len(queue) > 0 {
//         t := queue[0]
//         queue = queue[1:]

//         if t.i == n-1 && t.j == m-1 {
//             return t.step + 1
//         }

//         nexts := []Status{
//             {i:t.i-1,j:t.j},
//             {i:t.i+1,j:t.j},
//             {i:t.i,j:t.j-1},
//             {i:t.i,j:t.j+1},
//         }

//         for _, next := range nexts {
//             if next.i < 0 || next.i >= n || next.j < 0 || next.j >= m {
//                 continue
//             }

//             if maze[next.i][next.j] == 0 && !t.breakWall && !visit[next.i][next.j][0] {
//                 visit[next.i][next.j][0] = true
//                 next.step = t.step+1
//                 next.breakWall = t.breakWall
//                 queue = append(queue, next)
//             } else if maze[next.i][next.j] == 1 && !t.breakWall && !visit[next.i][next.j][1] {
//                 visit[next.i][next.j][1] = true
//                 next.breakWall = true
//                 next.step = t.step+1
//                 queue = append(queue, next)
//             } else if maze[next.i][next.j] == 0 && t.breakWall && !visit[next.i][next.j][1] {
//                 visit[next.i][next.j][1] = true
//                 next.step = t.step+1
//                 next.breakWall = t.breakWall
//                 queue = append(queue, next)
//             }
//         }
//     }
//     return -1
// }
