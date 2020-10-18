package main

func main() {

}

var (
	p1         = [5]int{1, 2, 3, 4, 5}
	p2         = [8]int{2, 1, 2, 3, 2, 4, 2, 5}
	p3         = [10]int{3, 3, 1, 1, 2, 2, 4, 4, 5, 5}
	a1, a2, a3 int
)

func solution(answers []int) []int {
	for i := 0; i < len(answers); i++ {
		if p1[i%len(p1)] == answers[i] {
			a1++
		}
		if p2[i%len(p2)] == answers[i] {
			a2++
		}
		if p3[i%len(p3)] == answers[i] {
			a3++
		}
	}
	max := maxInt(maxInt(a1, a2), a3)
	sol := make([]int, 0)
	if max == a1 {
		sol = append(sol, 1)
	}
	if max == a2 {
		sol = append(sol, 2)
	}
	if max == a3 {
		sol = append(sol, 3)
	}
	return sol
}
func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
