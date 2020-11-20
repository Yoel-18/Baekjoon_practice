package main

import (
	"fmt"
	"sort"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	for i := 0; i < b; i++ {
		for j := 0; j < a; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
func solution(arr []int, divisor int) []int {
	answer := make([]int, 0)
	for i := 0; i < len(arr); i++ {
		if arr[i]%divisor == 0 {
			answer = append(answer, arr[i])
		}
	}
	sort.Ints(answer)
	if len(answer) == 0 {
		answer = append(answer, -1)
	}
	return answer
}
