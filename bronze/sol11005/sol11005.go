package main

import "fmt"

func main() {
	var n, b int
	fmt.Scanf("%d %d", &n, &b)
	answer := make([]rune, 0)
	for n > 0 {
		if n%b < 10 {
			answer = append(answer, (rune)(n%b)+'0')
		} else {
			answer = append(answer, (rune)(n%b-10)+'A')
		}
		n /= b
	}
	for i := len(answer) - 1; i >= 0; i-- {
		fmt.Print(answer[i])
	}
}
