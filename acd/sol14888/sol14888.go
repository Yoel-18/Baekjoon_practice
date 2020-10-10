package main

import (
	"fmt"
	"math"
)

var (
	n   int
	a   []int
	op  = [4]int{}
	max = 0
	min = math.MaxInt64
)

func main() {

	fmt.Scan(&n)
	a = make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}
	for i := range op {
		fmt.Scan(&op[i])
	}

	findByDFS(a[0], 1)

	fmt.Println(max)
	fmt.Println(min)
}

func findByDFS(result, index int) {
	if index == n {
		if result > max {
			max = result
		}
		if min > result {
			min = result
		}
		return
	}

	for i := range op {
		if op[i] > 0 {
			op[i]--
			if i == 0 {
				findByDFS(result+a[index], index+1)
			}
			if i == 1 {
				findByDFS(result-a[index], index+1)
			}
			if i == 2 {
				findByDFS(result*a[index], index+1)
			}
			if i == 3 {
				findByDFS(result/a[index], index+1)
			}
			op[i]++
		}
	}
}
