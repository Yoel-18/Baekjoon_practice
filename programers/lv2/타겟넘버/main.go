package main

var (
	answer, size int
)

func main() {

}
func solution(numbers []int, target int) int {
	size = len(numbers)
	dfs(0, 0, target, numbers)
	return answer
}
func dfs(index, sum, target int, numbers []int) {
	if index == len(numbers) {
		if sum == target {
			answer++
		}
		return
	}
	dfs(index+1, sum+numbers[index], target, numbers)
	dfs(index+1, sum-numbers[index], target, numbers)
}
