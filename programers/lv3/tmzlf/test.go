package main

func main() {

}
func solution(n int, results [][]int) int {
	check := make([]int, n+1)
	for i := 1; i <= n; i++ {
		check[i] = n
	}
	for i := 0; i < n; i++ {
		if results[i][0] > results[i][1] {
			check[results[i][1]] = check[results[i][0]] - 1
		} else {
			check[results[i][0]] = check[results[i][1]] - 1
		}
	}
	return 0
}
