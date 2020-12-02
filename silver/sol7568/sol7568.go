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
	n := nextInt()
	user := make([]info, 0)
	for i := 0; i < n; i++ {
		a, b := nextInt(), nextInt()
		user = append(user, info{a, b})
	}
	for i := 0; i < n; i++ {
		rank := 1

		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			if user[i].weight < user[j].weight && user[i].height < user[j].height {
				rank++
			}
		}
		fmt.Print(rank, " ")
	}
}

type info struct {
	weight int
	height int
}

func nextInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}
func init() {
	sc.Split(bufio.ScanWords)
}
