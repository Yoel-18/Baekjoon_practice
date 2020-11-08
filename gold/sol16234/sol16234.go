package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	sc              = bufio.NewScanner(os.Stdin)
	n, l, r, answer int
	a               [][]int
	dx              = []int{-1, 1, 0, 0}
	dy              = []int{0, 0, -1, 1}
)

func main() {
	n, l, r = nextInt(), nextInt(), nextInt()
	a = make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, n)
		for j := 0; j < n; j++ {
			a[i][j] = nextInt()
		}
	}

	flag := true
	for flag {
		q := make([]int, 0)
		check := [10001]bool{}
		flag = false
		answer++

		for i := 0; i < n; i++ {
			if check[i] {
				continue
			}
			track := make([]int, 0)

			r := i / n
			c := i % n

			q = append(q, i)
			check[i] = true
			track = append(track, i)
			total := a[r][c]
			n = 1

			for len(q) > 0 {
				v := q[0]
				q = q[1:]
				r = v / n
				c = v % n

				for k := 0; k < 4; k++ {
					nx := c + dx[k]
					ny := r + dy[k]

					if nx >= 0 && nx < n && ny >= 0 && ny < n {
						next := ny*n + nx
						if check[next] {
							continue
						}

						count := a[ny][nx]
						pivot := a[r][c]
						if l <= absInt(count-pivot) && r >= absInt(count-pivot) {
							flag = true
							q = append(q, next)
							check[next] = true
							track = append(track, next)
							total += a[ny][nx]
							n++
						}
					}
				}
			}
			updateN := total / n
			if len(track) > 1 {
				for _, v := range track {
					r = v / n
					c = v % n
					a[r][c] = updateN
				}
			}
		}

	}
	fmt.Print(answer - 1)
}
func absInt(a int) int {
	if a < 0 {
		a = -1 * a
	}
	return a
}

func nextInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}
func init() {
	sc.Split(bufio.ScanWords)
}
