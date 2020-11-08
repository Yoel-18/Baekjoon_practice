package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// )

// var (
// 	sc           = bufio.NewScanner(os.Stdin)
// 	n, m, answer int
// 	vmap         [][]int
// )

// func main() {
// 	n, m = nextInt(), nextInt()
// 	vmap = make([][]int, n)
// 	for i := 0; i < n; i++ {
// 		vmap[i] = make([]int, m)
// 		for j := 0; j < m; j++ {
// 			vmap[i][j] = nextInt()
// 		}
// 	}
// 	find(0, 0, 0)
// 	fmt.Print(answer)
// }

// //	d: 벽 갯수
// func find(d, r, c int) {
// 	if d == 3 {
// 		var virus []point
// 		count := 0
// 		for i := range vmap {
// 			for j := range vmap[i] {
// 				if vmap[i][j] == 2 {
// 					var cur point
// 					cur.x, cur.y = i, j
// 					vmap[i][j] = 0
// 					virus = append(virus, cur)
// 				}
// 			}
// 		}
// 		for _, v := range virus {
// 			isVirus(v.x, v.y)
// 		}
// 		for i := range vmap {
// 			for j := range vmap[i] {
// 				if vmap[i][j] == 0 {
// 					count++
// 				}
// 			}
// 		}
// 		if answer < count {
// 			answer = count
// 		}
// 		return
// 	}
// 	for i := r; i < len(vmap); i++ {
// 		var j int
// 		if i == r {
// 			j = c
// 		} else {
// 			j = 0
// 		}
// 		for ; j < len(vmap[i]); j++ {
// 			newVmap := make([][]int, len(vmap))
// 			for ii := range newVmap {
// 				newVmap[ii] = append([]int{}, vmap[ii]...)
// 			}
// 			if newVmap[i][j] == 0 {
// 				newVmap[i][j] = 1
// 				find(newVmap, d+1, i, j)
// 			}
// 		}
// 	}
// }
// func isVirus(x, y int) {
// 	if vmap[x][y] == 0 {
// 		vmap[x][y] = 2
// 		if x+1 < len(vmap) {
// 			isVirus(x+1, y)
// 		}
// 		if y+1 < len(vmap) {
// 			isVirus(x+1, y)
// 		}
// 		if x-1 < len(vmap) {
// 			isVirus(x+1, y)
// 		}
// 		if y-1 < len(vmap) {
// 			isVirus(x+1, y)
// 		}
// 	}
// }

// type point struct {
// 	x, y int
// }

// func nextInt() int {
// 	sc.Scan()
// 	n, _ := strconv.Atoi(sc.Text())
// 	return n
// }
// func init() {
// 	sc.Split(bufio.ScanWords)
// }
