package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(solution([]int{1, 5, 2, 6, 3, 7, 4}, [][]int{{2, 5, 3}, {4, 4, 1}, {1, 7, 3}}))
}

func solution(array []int, commands [][]int) []int {
	answer := make([]int, 0)
	for i := 0; i < len(commands); i++ {
		a := make([]int, len(array))
		copy(a, array)
		cur := a[commands[i][0]-1 : commands[i][1]]
		sort.Ints(cur)
		answer = append(answer, cur[commands[i][2]-1])
	}
	return answer
}
