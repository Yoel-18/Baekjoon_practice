package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	sc = bufio.NewScanner(os.Stdin)
)

func main() {
	t := nextInt()
	for i := 0; i < t; i++ {
		n, m, count := nextInt(), nextInt(), 0
		q := make([]pri, 0)
		for j := 0; j < n; j++ {
			x := nextInt()
			q = append(q, pri{j, x})
		}
		for len(q) > 0 {
			cur := q[0]
			q = q[1:]
			check := true

			for _, que := range q {
				if que.priority > cur.priority {
					check = false
				}
			}
			if check {
				count++
				if cur.index == m {
					break
				}
			} else {
				q = append(q, cur)
			}
		}
		fmt.Println(count)
	}
}
func nextInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}
func init() {
	sc.Split(bufio.ScanWords)
}

type pri struct {
	index, priority int
}
