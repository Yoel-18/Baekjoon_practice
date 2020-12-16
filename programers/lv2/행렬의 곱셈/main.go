package main

func main() {

}
func solution(arr1 [][]int, arr2 [][]int) [][]int {
	answer := make([][]int, len(arr1))
	for i := 0; i < len(arr1); i++ {
		answer[i] = make([]int, len(arr2[0]))
		for j := 0; j < len(arr2[0]); j++ {
			sum := 0
			for k := 0; k < len(arr1[0]); k++ {
				sum += arr1[i][k] * arr2[k][j]
			}
			answer[i][j] = sum
		}

	}
	return answer
}
