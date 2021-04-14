package main

import "fmt"

func main() {
	fmt.Println(solution(5, []int{2, 4}, []int{1, 3, 5}))
	fmt.Println(solution(5, []int{2, 4}, []int{3}))
	fmt.Println(solution(2, []int{3}, []int{1}))
}
func solution(n int, lost []int, reserve []int) int {
	answer := n - len(lost)

	for i := 0; i < len(lost); i++ {
		for j := 0; j < len(reserve); j++ {
			if lost[i] == reserve[j] {
				answer++
				lost[i], reserve[j] = -1, -1
				break
			}
		}
	}

	for _, l := range lost {
		for i := 0; i < len(reserve); i++ {
			if l == reserve[i]-1 || l == reserve[i]+1 {
				answer++
				reserve[i] = -1
				break
			}
		}
	}

	return answer
}
