package main

import "fmt"

var (
	n, answer int
	q         [14]int
)

func main() {
	fmt.Scan(&n)
	find(0)
	fmt.Print(answer)
}

//	배열의 인덱스는 열
//	배열의 값은 행
//	첫번째 열에 행의 값을 찾고, 두번째 열에 또다른 행의 값을 찾고...
func find(depth int) {
	if depth == n {
		answer++
		return
	}
	for i := 0; i < n; i++ {
		q[depth] = i
		if possible(depth) {
			find(depth + 1)
		}
	}
}

//	놓을 수 있는 위치인가에 대한 판단
func possible(col int) bool {
	for i := 0; i < col; i++ {
		//	해당 열의 행과 i열의 행이 일치할 경우 -> 같은 행이 존재하는 경우
		if q[col] == q[i] {
			return false

			//	대각선상에 놓인 경우
			//	열의 차와 행의 차가 같을 경우이다
		} else if abs(col-i) == abs(q[col]-q[i]) {
			return false
		}
	}
	return true
}
func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}
