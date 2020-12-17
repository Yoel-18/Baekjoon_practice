package main

func main() {

}

//	두 숫자의 차이가 1일 때 만난다
// 한라운에는 참가자가 절반으로 줄어듬
// (현재번호 + 1) / 2 가 다음 라운드에서 자신의 번호
func solution(n int, a int, b int) int {
	answer := 1
	left, right := 0, 0
	if a > b {
		left = b
		right = a
	} else {
		left = a
		right = b
	}
	for {
		if left%2 != 0 && right-left == 1 {
			break
		}
		left = (left + 1) / 2
		right = (right + 1) / 2
		answer++
	}
	return answer
}
