package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	t := nextInt()
	for i := 0; i < t; i++ {
		x1, y1, r1, x2, y2, r2 := nextInt(), nextInt(), nextInt(), nextInt(), nextInt(), nextInt()
		if x1 == x2 && y1 == y2 {
			if r1 == r2 {
				fmt.Println(-1)
			} else {
				fmt.Println(0)
			}
			continue
		}
		rr1, rr2 := float64(r1), float64(r2)
		d := math.Sqrt(float64((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2)))
		if d == rr1+rr2 {
			fmt.Println(1)
		} else if d > rr1+rr2 {
			fmt.Println(0)
		} else if d+rr1 < rr2 || d+rr2 < rr1 {
			fmt.Println(0)
		} else if d+rr1 == rr2 || d+rr2 == rr1 {
			fmt.Println(1)
		} else {
			fmt.Println(2)
		}
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
