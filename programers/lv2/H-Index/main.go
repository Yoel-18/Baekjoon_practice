package main

import "sort"

func main() {

}
func solution(citations []int) int {
	sort.Slice(citations, func(i, j int) bool {
		return citations[i] > citations[j]
	})
	index := 0
	for index < len(citations) {
		if citations[index] <= index {
			break
		}
		index++
	}
	return index
}
