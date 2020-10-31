package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var (
	sc  = bufio.NewScanner(os.Stdin)
	num []int
)

func main() {
	n := nextInt()
	sum, max, min := 0, 0, math.MaxInt32
	num = make([]int, n)
	for i := 0; i < n; i++ {
		num[i] = nextInt()
		sum += num[i]
		if num[i] > max {
			max = num[i]
		}
		if num[i] < min {
			min = num[i]
		}
	}
	m := nextInt()
	if sum <= m {
		fmt.Print(max)
		return
	} else if min >= m {
		fmt.Println(min)
		return
	}
	left := m / n
	right := max
	result := 0
	for left <= right {
		mid := (left + right) / 2
		if mid == 0 {
			mid = 1
		}
		sum = 0
		for i := 0; i < n; i++ {
			sum += minInt(num[i], mid)
		}
		if sum > m {
			right = mid - 1
		} else {
			left = mid - 1
			if mid > max {
				result = mid
			}
		}
	}
	fmt.Print(result)
}
func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func minInt(a, b int) int {
	if a > b {
		return b
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
