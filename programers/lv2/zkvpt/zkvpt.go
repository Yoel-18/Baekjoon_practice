package main

func main() {

}
func solution(brown int, yellow int) []int {
	w, h := 0, 0

	for i := 1; i <= yellow/2+1; i++ {
		w = i
		if yellow%i == 0 {
			h = yellow / i
		} else {
			h = yellow/i + 1
		}

		if 2*w+2*h+4 == brown {
			break
		}
	}
	answer := []int{maxInt(w, h) + 2, minInt(w, h) + 2}
	return answer
}
func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func minInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}
