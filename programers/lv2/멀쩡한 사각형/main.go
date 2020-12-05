package main

func main() {

}
func solution(w int, h int) int {
	answer := 0
	answer = w*h - (w + h - gcd(w, h))
	return answer
}
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
