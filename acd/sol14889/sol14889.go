package main

import (
	"fmt"
	"math"
)

var (
	n      int
	stat   [20][20]int
	check  []bool
	answer = math.MaxInt64
)

func main() {
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Scan(&stat[i][j])
		}
	}
	check = make([]bool, n)

	makeTeam(0, 0)
	fmt.Println(answer)
}
func makeTeam(index, count int) {
	if count == n/2 {
		start := 0
		link := 0

		for i := 0; i < n; i++ {
			for j := i + 1; j < n; j++ {
				if check[i] != check[j] {
					continue
				}
				if check[i] {
					start += stat[i][j] + stat[j][i]
				} else {
					link += stat[i][j] + stat[j][i]
				}
			}
		}
		if answer > abs(start-link) {
			answer = abs(start - link)
		}
		return
	}
	for i := index; i < n; i++ {
		check[i] = true
		makeTeam(i+1, count+1)
		check[i] = false
	}
}
func abs(num int) int {
	if num > 0 {
		return num
	}
	return -num
}
