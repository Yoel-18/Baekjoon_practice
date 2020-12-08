package main

import "strconv"

func main() {

}

// return할 문자열의 길이가 4라면
// 맨 뒤의 3자리를 제외한 앞부분에서 제일 큰 수를 찾는다
// 이후 찾은 수부터 맨 뒤의 2자리를 제외한 수 중 큰 수를 찾는다
// 반복
func solution(number string, k int) string {
	answer := ""
	size := len(number)
	a, mm, nj := -1, -1, 0
	nn := size - k
	for i := 0; i < nn; i++ {
		for j := a + 1; j < size-(nn-i)+1; j++ {
			if mm < (int)(number[j]-'0') {
				mm = (int)(number[j] - '0')
				if mm == 9 {
					break
				}
				nj = j
			}
		}
		answer += strconv.Itoa(mm)
		a = nj
		mm = -1
	}
	return answer
}
