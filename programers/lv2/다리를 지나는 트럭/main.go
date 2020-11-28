package main

import "fmt"

func main() {
	fmt.Println(solution(2, 10, []int{7, 4, 5, 6}))
	fmt.Println(solution(100, 100, []int{10}))
	fmt.Println(solution(100, 100, []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10}))
}
func solution(bridgeLength int, weight int, truckWeights []int) int {
	answer := 0
	q := make([]int, 0)
	sum := 0
	for _, t := range truckWeights {
		for {
			if len(q) == 0 {
				q = append(q, t)
				sum += t
				answer++
				break
			} else if len(q) == bridgeLength {
				cur := q[0]
				q = q[1:]
				sum -= cur
			} else {
				if sum+t > weight {
					answer++
					q = append(q, 0)
				} else {
					q = append(q, t)
					sum += t
					answer++
					break
				}
			}
		}
	}
	return answer + bridgeLength
}
